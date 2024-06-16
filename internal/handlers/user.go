package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gustafer/linkord/internal/database"
	"github.com/gustafer/linkord/internal/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func PrivateUserInfo(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"message": "protected!",
	}

	b, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetUserId(w, r)
	user, err := database.GetUser(userId)
	if err != nil {
		http.Error(w, "could not get user", http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
