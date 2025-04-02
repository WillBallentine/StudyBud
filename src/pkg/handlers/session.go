package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		logrus.Infof("session retrival err was not nil: %v", err)
		return AuthError
	}

	email, ok := session.Values["email"].(string)
	if !ok || email == "" {
		logrus.Info("user email missing from session")
		return AuthError
	}

	ctx, ctxErr := context.WithTimeout(context.TODO(), time.Duration(config.App.Timeout)*time.Second)
	defer ctxErr()

	user, exists := mongo_repo.GetUserByEmail(email, ctx)
	if !exists {
		logrus.Info("user not returning")
		return AuthError
	}

	st, err := r.Cookie("session_token")
	logrus.Infof("header token: %v", st)
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		logrus.Infof("db token: %v", user.SessionToken)
		logrus.Info("session token error")
		return AuthError
	}

	logrus.Info("authorized")
	return nil

}
