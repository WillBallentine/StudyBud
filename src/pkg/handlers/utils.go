package handlers

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"studybud/src/cmd/utils"
	model "studybud/src/pkg/models"
	"studybud/src/pkg/mongodb"
	service "studybud/src/pkg/service"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Templates *template.Template
var mongo_plan_repo = mongodb.Initialize(config, "plans")
var mongo_user_repo = mongodb.Initialize(config, "users")

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

	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}

	if r.Method == http.MethodPost {

		csrfToken := r.FormValue("csrf_token")

		session, _ := store.Get(r, "session-name")
		expectedToken, _ := session.Values["csrf_token"].(string)

		if csrfToken == "" || csrfToken != expectedToken {
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

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
		processPlanChan := make(chan model.StudyPlan)
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

		var syllabus model.SyllabusData
		select {
		case result := <-processResultChan:
			syllabus = result
			userId, _ := primitive.ObjectIDFromHex(session.Values["userId"].(string))
			ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
			defer ctxErr()

			mongo_user_repo.UpsertSyllabusId(userId, syllabus.ID, ctx)
		case err := <-errChan:
			logrus.Info("err chan")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					logrus.Errorf("recovered from panic in syllabus parsing: %v", r)
					errChan <- fmt.Errorf("internal error in syllabus parsing")
				}
			}()
			service.ProcessStudyPlan(syllabus, processPlanChan, errChan)
		}()

		var studyPlan model.StudyPlan
		select {
		case result := <-processPlanChan:
			studyPlan = result
			userId, _ := primitive.ObjectIDFromHex(session.Values["userId"].(string))
			ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
			defer ctxErr()

			mongo_user_repo.UpsertStudyPlanId(userId, studyPlan.ID, ctx)
		case err := <-errChan:
			logrus.Info("studyplan error chan")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmple, err := template.ParseFiles("web/templates/pages/confirm.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			CSRFToken string
			StudyPlan model.StudyPlan
		}{
			CSRFToken: csrfToken,
			StudyPlan: studyPlan,
		}

		tmple.Execute(w, data)
		return
	}

	session, _ := store.Get(r, "session-name")
	csrfToken, _ := session.Values["csrf_token"].(string)
	if csrfToken == "" {
		csrfToken = utils.GenerateToken(32)
		session.Values["csrf_token"] = csrfToken
		session.Save(r, w)
	}

	tmpl, err := template.ParseFiles("web/templates/pages/upload.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]string{"CSRFToken": csrfToken})
}

func EditStudyPlanHandler(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	csrfToken := r.FormValue("csrf_token")
	session, _ := store.Get(r, "session-name")
	expectedToken, _ := session.Values["csrf_token"].(string)
	if csrfToken != expectedToken {
		http.Error(w, "CSRF token mismatch", http.StatusForbidden)
		return
	}

	planID := r.FormValue("plan_id")

	re := regexp.MustCompile(`^ObjectID\("(.*)"\)$`)
	matches := re.FindStringSubmatch(planID)

	if len(matches) != 2 {
		http.Error(w, "Invalid plan ID format", http.StatusBadRequest)
		return
	}

	cleanedPlanID := matches[1]

	objID, err := primitive.ObjectIDFromHex(cleanedPlanID)
	if err != nil {
		logrus.Errorf("Error converting planId to ObjectID: %v", err)
		http.Error(w, "Invalid plan ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.App.Timeout)*time.Second)
	defer cancel()

	plan, err := mongo_plan_repo.GetStudyPlanByID(objID, ctx)
	if err != nil {
		http.Error(w, "Study plan not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/pages/edit_studyplan.html")
	if err != nil {
		http.Error(w, "Template parse error", http.StatusInternalServerError)
		return
	}

	data := struct {
		CSRFToken string
		StudyPlan *model.StudyPlan
	}{
		CSRFToken: csrfToken,
		StudyPlan: plan,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		logrus.Errorf("error executing template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func SaveStudyPlanHandler(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	csrfToken := r.FormValue("csrf_token")
	session, _ := store.Get(r, "session-name")
	expectedToken, _ := session.Values["csrf_token"].(string)
	if csrfToken != expectedToken {
		http.Error(w, "CSRF token mismatch", http.StatusForbidden)
		return
	}

	planID := r.FormValue("plan_id")
	objID, err := primitive.ObjectIDFromHex(planID)
	logrus.Infof("planId: %v", planID)
	logrus.Infof("objId: %v", objID)
	if err != nil {
		http.Error(w, "Invalid plan ID", http.StatusBadRequest)
		return
	}
	// You would parse all fields from the form and update them into the StudyPlan object.
	// Here’s a simplified example:
	//updatedPlan := model.StudyPlan{
	//	ID: objID,
	// You'd fill in all fields based on form values
	// StudyBlocks: ...
	//}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.App.Timeout)*time.Second)
	//defer cancel()

	//err = mongo_repo.UpdateStudyPlan(updatedPlan, ctx)
	//if err != nil {
	//	http.Error(w, "Failed to update study plan", http.StatusInternalServerError)
	//return
	//}

	http.Redirect(w, r, "/upload", http.StatusSeeOther)
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
