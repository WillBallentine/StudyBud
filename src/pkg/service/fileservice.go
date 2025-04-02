package service

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"studybud/src/cmd/utils"
	"studybud/src/pkg/entity"
	"studybud/src/pkg/models"
	"studybud/src/pkg/mongodb"

	"github.com/go-resty/resty/v2"
	"github.com/ledongthuc/pdf"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const openaiAPIURL = "https://api.openai.com/v1/chat/completions"

var config = utils.Read_Configuration(utils.Read())
var mongo_syll_repo = mongodb.Initialize(config, "syllabus")
var mongo_plan_repo = mongodb.Initialize(config, "plans")

func ExtractPdfTextAsync(pdfPath string, textChan chan<- string, errChan chan<- error) {

	f, err := os.Open(pdfPath)
	if err != nil {
		errChan <- fmt.Errorf("error opening PDF: %v", err)
		logrus.Errorf("cannot open pdf: %v", err)
		return
	}

	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		errChan <- fmt.Errorf("error getting file info: %v", err)
		logrus.Errorf("error getting file info: %v", err)
		return
	}

	logrus.Infof("processing pdf: %s", fileInfo.Name())

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		errChan <- fmt.Errorf("error resetting pointer: %v", err)
		logrus.Errorf("error resetting pointer: %v", err)
		return
	}

	r, err := pdf.NewReader(f, fileInfo.Size())
	if err != nil {
		errChan <- fmt.Errorf("error reading PDF: %v", err)
		logrus.Errorf("cannot read PDF: %v", err)
		return
	}

	var extractedText bytes.Buffer
	for i := 1; i <= r.NumPage(); i++ {
		logrus.Infof("processing page %d", i)
		page := r.Page(i)
		if page.V.IsNull() {
			logrus.Errorf("error: page %d is empty", i)

		}
		text := page.V.Text()
		logrus.Infof("extracted text from page %d: %s", i, text[:min(len(text), 100)])
		extractedText.WriteString(page.V.Text() + "\n")
	}

	logrus.Info("finished processing pages")
	if extractedText.Len() == 0 {
		errChan <- fmt.Errorf("no text extracted from pdf")
		logrus.Error("no text extracted")
		return
	}

	textChan <- extractedText.String()
}

func ExtractDocxTextAsync(path string, textChan chan<- string, errChan chan<- error) {
	docxFile, err := os.Open(path)
	if err != nil {
		errChan <- fmt.Errorf("error opening docx file: %v", err)
		logrus.Errorf("error opening docx file: %v", err)
		return
	}

	defer docxFile.Close()

	tmpDir, err := ioutil.TempDir("", "docx_extract")
	if err != nil {
		errChan <- fmt.Errorf("error creating temp dir: %v", err)
		logrus.Errorf("error creating temp dir: %v", err)
		return
	}

	defer os.RemoveAll(tmpDir)

	tempPath := fmt.Sprintf("%s/temp.docx", tmpDir)
	outfile, err := os.Create(tempPath)
	if err != nil {
		errChan <- fmt.Errorf("error creating temp file: %v", err)
		logrus.Errorf("error creating temp file: %v", err)
		return
	}

	defer outfile.Close()

	_, err = io.Copy(outfile, docxFile)
	if err != nil {
		errChan <- fmt.Errorf("error copying file: %v", err)
		logrus.Errorf("error copying file: %v", err)
		return
	}

	err = unzip(tempPath, tmpDir)
	if err != nil {
		errChan <- fmt.Errorf("error uzipping file: %v", err)
		logrus.Errorf("error unzipping file: %v", err)
		return
	}

	documentXML := fmt.Sprintf("%s/word/document.xml", tmpDir)

	xmlContent, err := ioutil.ReadFile(documentXML)
	if err != nil {
		errChan <- fmt.Errorf("error reading document.xml: %v", err)
		logrus.Errorf("error reading document.xml: %v", err)
		return
	}

	re := regexp.MustCompile("<[^>]*>")
	text := re.ReplaceAllString(string(xmlContent), "")
	text = xmlEscape(text)

	textChan <- text

}

func ParseSyllabusWithOpenAI(text string, resultChan chan<- model.SyllabusData, errChan chan<- error) {
	client := resty.New()

	prompt := fmt.Sprintf(`
		Extract the following information from the given college course syllabus text:
		- Professor Name
		- Course Title
		- Course Code
		- Assignments (with due dates and description if available. if there are any related texts for the assignment, add their title if available)
		- Required Textbooks (titles & authors. please include a link if available)
		- School
		- Semester
		
		Return the response strictly in this JSON format:
		{
			"professor_name": "",
			"course_title": "",
			"course_code": "",
			"assignments": [{"title": "", "desc": "", "due_date": "", "related_texts": [{"title": ""}]}],
			"required_texts": [{"title": "", "author": "", "link": ""}],
			"school": "",
			"semester": ""
		}

		Text:
		"""%s"""
		Respond only with valid JSON
	`, text)

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY")).SetBody(map[string]interface{}{
		"model":    "gpt-4o",
		"messages": []map[string]string{{"role": "user", "content": prompt}}, "temperature": 0.2,
	}).Post(openaiAPIURL)

	if err != nil {
		errChan <- fmt.Errorf("error sending to openai: %v", err)
	}

	var openAIResponseData, unpackErr = extractOpenAIResponse(resp, err)

	if unpackErr != nil {
		errChan <- fmt.Errorf("error decoding reponse from openai: %v", unpackErr)
	}

	var syllabusData model.SyllabusData
	err = json.Unmarshal([]byte(openAIResponseData), &syllabusData)
	if err != nil {
		errChan <- fmt.Errorf("error unmarshaling into syllabusdata model: %v", err)
		logrus.Errorf("error unmarshaling into syllabusdata model: %v", err)
		return
	}

	var oId = saveSyllabus(syllabusData)
	syllabusData.ID = oId

	resultChan <- syllabusData
}

func ProcessStudyPlan(syllabusData model.SyllabusData, resultChan chan<- model.StudyPlan, errChan chan<- error) {
	client := resty.New()

	syllabusJson, _ := json.Marshal(syllabusData)
	prompt := fmt.Sprintf(`
	You are an AI that generates a structured study plan based on syllabus data. The study plan should be returned as a JSON object with the following structure:

{
    "course_name": "Course Name",
    "course_code": "Course Code",
    "semester": "Semester",
    "study_blocks": [
        {
            "title": "Study Block Title",
            "description": "Description of the study block",
            "start_date": "YYYY-MM-DD", // The date when the study block starts
            "due_date": "YYYY-MM-DD",   // The due date for this block (e.g., assignment due date)
            "priority": 1,              // Priority (1 = high, 2 = medium, 3 = low)
            "estimated_time": "2h",     // Estimated time for completion (e.g., 2 hours)
            "completed": false,         // Whether the task is completed (false initially)
            "related_texts": [
                {
                    "title": "Reference Title",
                    "author": "Author Name",
                    "link": "URL to reference"
                }
            ]
        }
    ]
}

Each study block should be linked to the corresponding assignments in the syllabus if applicable, with due dates, estimated time to complete, and any related references or texts provided in the syllabus.

Please ensure the generated study plan covers the following:
1. Title and description of the study blocks.
2. Start date, due date, and estimated time for each block. if no dates are provided in the data given you, do not create dates. leave those fields empty.
3. Prioritization of tasks (using a scale of 1 to 3).
4. Associated references (textbooks or other materials) for each block.

Example:

{
    "course_name": "Introduction to AI",
    "course_code": "CS101",
    "semester": "Spring 2025",
    "study_blocks": [
        {
            "title": "Read Chapter 1",
            "description": "Study the basics of AI, focusing on definitions and history.",
            "start_date": "2025-04-01",
            "due_date": "2025-04-05",
            "priority": 1,
            "estimated_time": "3h",
            "completed": false,
            "related_texts": [
                {
                    "title": "Artificial Intelligence: A Modern Approach",
                    "author": "Stuart Russell and Peter Norvig",
                    "link": "http://link_to_textbook.com"
                }
            ]
        }
    ]
}

Please use this format for all study blocks.ata, prepare for me a studyplan.

		Respond only with valid JSON

		syllabus data: %s
		`, string(syllabusJson))

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetHeader("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY")).SetBody(map[string]interface{}{
		"model":    "gpt-4o",
		"messages": []map[string]string{{"role": "user", "content": prompt}}, "temperature": 0.2,
	}).Post(openaiAPIURL)

	if err != nil {
		errChan <- fmt.Errorf("error sending to openai: %v", err)
		logrus.Errorf("error sending to plan to openai: %v", err)
	}

	logrus.Infof("study plan response: %v", resp)
	var openAIResponseData, unpackErr = extractOpenAIResponse(resp, err)

	if unpackErr != nil {
		errChan <- fmt.Errorf("error decoding reponse from openai: %v", unpackErr)
		logrus.Errorf("error sending to plan to openai: %v", err)
	}

	var studyPlan model.StudyPlan
	err = json.Unmarshal([]byte(openAIResponseData), &studyPlan)
	if err != nil {
		errChan <- fmt.Errorf("error unmarshaling into studyplan model: %v", err)
		logrus.Errorf("error unmarshaling into studyplan model: %v", err)
		return
	}

	var oId = saveStudyPlan(studyPlan)
	studyPlan.ID = oId

	resultChan <- studyPlan
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outfile, err := os.Create(fpath)

		if err != nil {
			return err
		}

		defer outfile.Close()

		rc, err := f.Open()
		if err != nil {
			return err
		}

		defer rc.Close()

		_, err = io.Copy(outfile, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func xmlEscape(text string) string {
	r := strings.NewReplacer(
		"&lt;", "<",
		"&gt;", ">",
		"&amp;", "&",
		"&quot;", `"`,
		"&apos;", "'",
	)

	return r.Replace(text)
}

func extractJSON(response string) string {
	if strings.HasPrefix(response, "```json") {
		response = strings.TrimPrefix(response, "```json")
	}

	if strings.HasSuffix(response, "```") {
		response = strings.TrimSuffix(response, "```")
	}

	return strings.TrimSpace(response)
}

func saveSyllabus(syll model.SyllabusData) primitive.ObjectID {
	ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
	defer ctxErr()

	var newsyllabusEntity *entity.SyllabusDataEntity
	newsyllabusEntity = &entity.SyllabusDataEntity{
		Instructor:    syll.Instuctor,
		CourseCode:    syll.CourseCode,
		CourseName:    syll.CourseName,
		Assignments:   toAssignmentEntity(syll.Assignments),
		RequiredTexts: toRequiredTextsEntity(syll.References),
		School:        syll.School,
		Semester:      syll.Semester,
	}

	oId, err := mongo_syll_repo.AddSyllabus(*newsyllabusEntity, ctx)

	if err != nil {
		logrus.Error("failed to save syllabus to db")
	}

	return oId
}

func saveStudyPlan(plan model.StudyPlan) primitive.ObjectID {
	ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
	defer ctxErr()

	oId, err := mongo_plan_repo.AddStudyPlan(plan, ctx)

	if err != nil {
		logrus.Error("failed to save studyplan to db")
	}

	return oId
}

func toAssignmentEntity(input []model.Assignment) []entity.AssignmentEntity {
	var entities []entity.AssignmentEntity
	for _, assignment := range input {
		entities = append(entities, entity.AssignmentEntity{
			Title:       assignment.Title,
			Description: assignment.Description,
			DueDate:     assignment.DueDate,
			Refs:        toRequiredTextsEntity(assignment.Refs),
		})
	}
	return entities
}

func toRequiredTextsEntity(input []model.Reference) []entity.ReferenceEntity {
	var entities []entity.ReferenceEntity
	for _, ref := range input {
		entities = append(entities, entity.ReferenceEntity{
			Title:  ref.Title,
			Author: ref.Author,
			Link:   ref.Link,
		})
	}

	return entities
}

func extractOpenAIResponse(resp *resty.Response, err error) (string, error) {

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		logrus.Errorf("error unmarshalling response from openai: %v", err)
		return "", err
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		logrus.Errorf("unexpected API response format")
		return "", err
	}

	message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok || len(choices) == 0 {
		logrus.Errorf("unexpected API response format")
		return "", err
	}

	content, ok := message["content"].(string)
	if !ok || len(choices) == 0 {
		logrus.Errorf("unexpected API response format")
		return "", err
	}

	cleanedjson := extractJSON(content)
	logrus.Infof("cleaned json: %v", cleanedjson)

	return cleanedjson, nil
}
