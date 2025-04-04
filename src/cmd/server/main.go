package main

import (
	"fmt"
	"log"
	"net/http"
	"studybud/src/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

var router *chi.Mux

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static/favicon.ico")
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error
	handlers.Catch(err)

	handlers.LoadTemplates()

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/edit-studyplan", handlers.EditStudyPlanHandler)
	http.HandleFunc("/save-studyplan", handlers.SaveStudyPlanHandler)
	http.HandleFunc("/", handlers.HomeHandler)

	port := ":8080"
	fmt.Println("server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
