package service

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"studybud/src/pkg/models"

	"github.com/jdkato/prose/v2"
	"github.com/ledongthuc/pdf"
	"github.com/sirupsen/logrus"
)

func ExtractPdfTextAsync(pdfPath string, textChan chan<- string, errChan chan<- error) {
	defer close(textChan)
	defer close(errChan)

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

	r, err := pdf.NewReader(f, fileInfo.Size())
	if err != nil {
		errChan <- fmt.Errorf("error reading PDF: %v", err)
		logrus.Errorf("cannot read PDF: %v", err)
		return
	}

	var extractedText bytes.Buffer
	for i := 1; i <= r.NumPage(); i++ {
		page := r.Page(i)
		extractedText.WriteString(page.V.Text() + "\n")
	}

	textChan <- extractedText.String()
}

func ParseSyllabus(text string, resultChan chan<- model.SyllabusData, errChan chan<- error) (model.SyllabusData, error) {
	var syllabus model.SyllabusData

	doc, err := prose.NewDocument(text)
	if err != nil {
		logrus.Errorf("failed to process pdf: %v", err)
		errChan <- fmt.Errorf("failed to process pdf: %v", err)
		return model.SyllabusData{}, err
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

	return syllabus, nil

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
