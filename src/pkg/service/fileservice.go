package service

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"studybud/src/pkg/models"

	"github.com/go-resty/resty/v2"
	"github.com/ledongthuc/pdf"
	"github.com/sirupsen/logrus"
)

const openaiAPIURL = "https://api.openai.com/v1/chat/completions"

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

	logrus.Infof("resp: %v", resp)

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		errChan <- fmt.Errorf("could not unmarshal response from openai: %v", err)
		logrus.Errorf("error unmarshalling response from openai: %v", err)
		return
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		errChan <- fmt.Errorf("unexpected API response format")
		logrus.Errorf("unexpected API response format")
		return
	}

	message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok || len(choices) == 0 {
		errChan <- fmt.Errorf("unexpected API response format")
		logrus.Errorf("unexpected API response format")
		return
	}

	content, ok := message["content"].(string)
	if !ok || len(choices) == 0 {
		errChan <- fmt.Errorf("unexpected API response format")
		logrus.Errorf("unexpected API response format")
		return
	}

	cleanedjson := extractJSON(content)

	logrus.Infof("cleaned json: %v", cleanedjson)

	var syllabusData model.SyllabusData
	err = json.Unmarshal([]byte(cleanedjson), &syllabusData)
	if err != nil {
		errChan <- fmt.Errorf("error unmarshaling into syllabusdata model: %v", err)
		logrus.Errorf("error unmarshaling into syllabusdata model: %v", err)
		return
	}

	resultChan <- syllabusData
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

func jsonStartIndex(response string) int {
	start := -1
	for i, ch := range response {
		if ch == '{' {
			start = i
			break
		}
	}
	return start
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
