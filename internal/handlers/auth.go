package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/gustafer/linkord/configs"
	"github.com/gustafer/linkord/internal/database"
	"github.com/markbates/goth/gothic"
)

var (
	authCookie string = "Auth"
)

func newCookie(value string) http.Cookie {
	return http.Cookie{
		Name:     authCookie,
		SameSite: http.SameSiteNoneMode,
		Secure:   false,
		Value:    value,
		HttpOnly: true,
		Path:     "/",
	}
}

func SignAuthCookie(userId string) (token string, err error) {
	claims := jwt.MapClaims{
		"user_id": userId,
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := newToken.SignedString([]byte(configs.LoadAuthKey()))
	if err != nil {
		return "", err
	}
	return signedToken, err
}

func GetAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	// complete user auth
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// upsert user to db
	if err = database.UpsertUser(user.UserID, user.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := SignAuthCookie(user.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := newCookie(token)
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, fmt.Sprintf("http://%v/", configs.GetRedirectUrl()), http.StatusFound)
}

func GetLogout(w http.ResponseWriter, r *http.Request) {
	err := gothic.Logout(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deleteAuthCookie := &http.Cookie{
		Name:     authCookie,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, deleteAuthCookie)

	http.Redirect(w, r, fmt.Sprintf("http://%v/", configs.GetRedirectUrl()), http.StatusTemporaryRedirect)
}

func GetAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		if err := database.UpsertUser(user.UserID, user.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token, err := SignAuthCookie(user.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cookie := newCookie(token)
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, fmt.Sprintf("http://%v/", configs.GetRedirectUrl()), http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}
