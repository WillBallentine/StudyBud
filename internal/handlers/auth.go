package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var templates = template.Must(template.ParseGlob("web/templates/**/*.html"))
var store = sessions.NewCookieStore([]byte("some_key"))

// faking a user db for testing
var users = map[string]string{}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := Templates.ExecuteTemplate(w, "base.html", map[string]string{"Title": "Register"})
		if err != nil {
			fmt.Println("register template failed to execute")
		}
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Save user
	users[email] = string(hashedPassword)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		// Validate user
		hashedPassword, exists := users[email]
		if !exists || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
			//eventually some notification to user and retry logic. dont want to take to registration just for typos
			http.Redirect(w, r, "/register", http.StatusSeeOther)
		}

		// Create session
		session, _ := store.Get(r, "session")
		session.Values["email"] = email
		session.Save(r, w)

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles("web/templates/pages/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "email")
	session.Save(r, w)
}
