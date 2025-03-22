package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Templates *template.Template

func LoadTemplates() {
	var err error
	Templates, err = template.ParseGlob("src/web/templates/**/*.html")
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
	t, err := template.ParseFiles("src/web/templates/pages/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
