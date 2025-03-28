package service

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"studybud/src/pkg/models"

	"github.com/jdkato/prose/v2"
	"github.com/ledongthuc/pdf"
	"github.com/otiai10/copy"
	"github.com/sirupsen/logrus"
)

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

	//TODO: this is where the docx is erroring out
	err = copy.Copy(path, tmpDir)
	if err != nil {
		errChan <- fmt.Errorf("error unzipping docx file: %v", err)
		logrus.Errorf("error unzipping docx file: %v", err)
		return
	}

	documentXML := fmt.Sprintf("%s/word/document.xml", tmpDir)

	xmlContent, err := ioutil.ReadFile(documentXML)
	if err != nil {
		errChan <- fmt.Errorf("error reading document.xml: %v", err)
		logrus.Errorf("error reading document.xml: %v", err)
		return
	}

	text := strings.ReplaceAll(string(xmlContent), "<w:t>", "")
	text = strings.ReplaceAll(text, "</w:t>", "")

	textChan <- text

}

func ParseSyllabus(text string, resultChan chan<- model.SyllabusData, errChan chan<- error) {
	var syllabus model.SyllabusData

	doc, err := prose.NewDocument(text)
	if err != nil {
		logrus.Errorf("failed to process pdf: %v", err)
		errChan <- fmt.Errorf("failed to process pdf: %v", err)
	}

	for _, ent := range doc.Entities() {
		if strings.Contains(strings.ToLower(ent.Label), "person") {
			if strings.Contains(strings.ToLower(ent.Text), "prof") || strings.Contains(strings.ToLower(ent.Text), "instructor") {
				syllabus.Instuctor = ent.Text
				break
			}
		}
	}

	for _, ent := range doc.Entities() {
		if strings.Contains(strings.ToLower(ent.Text), "assignment") {
			syllabus.Assignments = append(syllabus.Assignments, ent.Text)
		}
	}

	for _, ent := range doc.Entities() {
		if strings.Contains(strings.ToLower(ent.Text), "books") || strings.Contains(strings.ToLower(ent.Text), "texts") {
			syllabus.References = append(syllabus.References, ent.Text)
		}
	}

	syllabus.CourseName = "temp. eventually this will be pulled from user input when uploading syllabus"
	syllabus.School = "temp eventually will be pulled from user account data"
	syllabus.Semester = "temp. eventually pulled from user account data"

	resultChan <- syllabus

}

func extractSection(text string, startKeyword string, endKeyword string) string {
	startIdx := strings.Index(strings.ToLower(text), strings.ToLower(startKeyword))
	if startIdx == -1 {
		return ""
	}

	remainingText := text[startIdx+len(startKeyword):]
	endIdx := strings.Index(strings.ToLower(remainingText), strings.ToLower(endKeyword))

	if endIdx != -1 {
		return remainingText[:endIdx]
	}

	return remainingText
}
