package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"studybud/src/cmd/utils"
	"studybud/src/pkg/entity"
	"studybud/src/pkg/mongodb"

	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var config = utils.Read_Configuration(utils.Read())
var mongo_repo = mongodb.Initialize(config)

var templates = template.Must(template.ParseGlob("web/templates/**/*.html"))
var store = sessions.NewCookieStore([]byte("some_key"))

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

	var newUser *entity.User

	newUser = &entity.User{
		FirstName:         "test",
		LastName:          "test",
		Email:             email,
		Password:          string(hashedPassword),
		School:            "sbts",
		SubscriptionLevel: "premium",
	}

	// Register user
	ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
	defer ctxErr()

	oId, err := mongo_repo.AddUser(*newUser, ctx)

	if err != nil {
		logrus.Info("failed to register user")
		http.Redirect(w, r, "/register", http.StatusBadRequest)
	}

	session, _ := store.Get(r, "session")
	session.Values["userId"] = oId
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
		defer ctxErr()
		// Validate user
		hashedPassword, exists := mongo_repo.GetUserByEmail(email, ctx)
		if !exists || bcrypt.CompareHashAndPassword([]byte(hashedPassword.Password), []byte(password)) != nil {
			//eventually some notification to user and retry logic. dont want to take to registration just for typos
			http.Redirect(w, r, "/register", http.StatusSeeOther)
		}

		// Create session
		session, _ := store.Get(r, "session")
		session.Values["email"] = email
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles("web/templates/pages/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// this needs a full fleshing out. currently just clearing session data.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "userId")
	session.Save(r, w)
}
