package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gustafer/linkord/internal/database"
)

func ProjectInfo(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{
		"message": "Go API that powers Linkord.",
	}
	b, _ := json.Marshal(message)
	w.Write(b)
}

func Health(w http.ResponseWriter, r *http.Request) {
	db := database.NewConn()

	if err := db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message": "health ok!",
	}
	b, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
