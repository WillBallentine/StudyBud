package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	model "studybud/src/pkg/models"
	service "studybud/src/pkg/service"

	"github.com/sirupsen/logrus"
)

var Templates *template.Template

func LoadTemplates() {
	var err error
	Templates, err = template.ParseGlob("web/templates/**/*.html")
	if err != nil {
		log.Fatalf("error loading templates %v", err)
	}

	fmt.Println("templates loaded")
}

func Catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("web/templates/pages/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		err := r.ParseMultipartForm(10 >> 20)
		if err != nil {
			http.Error(w, "error parsing form data", http.StatusBadRequest)
			logrus.Errorf("error parsing form data: %v", err)
		}

		file, _, err := r.FormFile("file")

		if err != nil {
			http.Error(w, "failed to read file", http.StatusBadRequest)
			logrus.Errorf("failed to read file: %v", err)
			return
		}

		defer file.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, file)
		if err != nil {
			http.Error(w, "failed to buffer file", http.StatusInternalServerError)
			logrus.Errorf("failed to buffer file: %v", err)
			return
		}

		fileReader := bytes.NewReader(buf.Bytes())

		textChan := make(chan string)
		processResultChan := make(chan model.SyllabusData)
		fileTypeChan := make(chan string)
		errChan := make(chan error)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					logrus.Errorf("recovered from panic in getFileType: %v", r)
					errChan <- fmt.Errorf("internal error in file type extraction")
				}
			}()
			getFileType(fileReader, fileTypeChan, errChan)
		}()

		var fileType string
		select {
		case filetype := <-fileTypeChan:
			fileType = filetype
		case err := <-errChan:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileName := fmt.Sprintf("syllabus-*.%s", fileType)
		fileReader.Seek(0, io.SeekStart)

		tempfile, err := os.CreateTemp("", fileName)
		if err != nil {
			http.Error(w, "could not create temp file", http.StatusInternalServerError)
			logrus.Errorf("could not create temp file: %v", err)
			return
		}

		defer tempfile.Close()

		_, err = io.Copy(tempfile, fileReader)
		if err != nil {
			http.Error(w, "could not save file", http.StatusInternalServerError)
			logrus.Errorf("could not save file: %v", err)
			return
		}

		logrus.Infof("file successfully written: %s", tempfile.Name())

		if fileType == "pdf" {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						logrus.Errorf("recovered from panic: %v", r)
						errChan <- fmt.Errorf("internal error in pdf extraction")
					}
				}()
				service.ExtractPdfTextAsync(tempfile.Name(), textChan, errChan)
			}()

		} else if fileType == "docx" {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						logrus.Errorf("recovered from panic: %v", r)
						errChan <- fmt.Errorf("internal error in pdf extraction")
					}
				}()
				service.ExtractDocxTextAsync(tempfile.Name(), textChan, errChan)
			}()
		}

		var extractedText string
		select {
		case text := <-textChan:
			extractedText = text
		case err := <-errChan:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if extractedText == "" {
			http.Error(w, "failed to extract text", http.StatusInternalServerError)
			logrus.Errorf("failed to extract text: %v", err)
			return
		}

		os.Remove(tempfile.Name())

		go func() {
			defer func() {
				if r := recover(); r != nil {
					logrus.Errorf("recovered from panic in syllabus parsing: %v", r)
					errChan <- fmt.Errorf("internal error in syllabus parsing")
				}
			}()
			service.ParseSyllabusWithOpenAI(extractedText, processResultChan, errChan)
		}()

		var response model.SyllabusData
		select {
		case result := <-processResultChan:
			logrus.Info("result chan")
			response = result
		case err := <-errChan:
			logrus.Info("err chan")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logrus.Infof("some response %v", response)

		//TODO: need to rework this. just for quick testing purposes.


		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	tmpl, err := template.ParseFiles("web/templates/pages/upload.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func getFileType(file io.Reader, fileTypeChan chan<- string, errChan chan<- error) {
	var buf [512]byte
	n, err := file.Read(buf[:])

	if err != nil {
		errChan <- fmt.Errorf("error reading file in type detection: %v", err)
		logrus.Errorf("could not read file to extract file type")
		return
	}

	fileType := http.DetectContentType(buf[:n])

	logrus.Infof("file type: %s", fileType)

	if strings.Contains(fileType, "pdf") {
		fileTypeChan <- "pdf"
		return
	} else if strings.Contains(fileType, "word") || strings.Contains(fileType, "zip") {
		fileTypeChan <- "docx"
		return
	}

	errChan <- fmt.Errorf("unsupported file type: %s", fileType)

}
