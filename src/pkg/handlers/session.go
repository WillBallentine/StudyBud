package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return AuthError
	}

	userId := session.Values["userid"]
	logrus.Info(userId)
	ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
	defer ctxErr()

	user, err := mongo_repo.GetUserById(userId.(primitive.ObjectID), ctx)
	if err != nil {
		logrus.Info("user not returning")
		return AuthError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		logrus.Info("session token error")
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		logrus.Info("csrf token error")
		return AuthError
	}

	return nil

}
