package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var templates = template.Must(template.ParseGlob("web/templates/**/*.html"))
var store = sessions.NewCookieStore([]byte("some_key"))
var users = map[string]string{}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "base.html", map[string]string{"Title": "Register"})
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
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "base.html", map[string]string{"Title": "Login"})
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validate user
	hashedPassword, exists := users[email]
	if !exists || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create session
	session, _ := store.Get(r, "session")
	session.Values["email"] = email
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "email")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "base.html", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
