package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gustafer/linkord/internal/database"
)

var game *database.Game

func CreateGame(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	validate := validator.New()
	if err := validate.Struct(game); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdGameId, err := database.CreateGame(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := map[string]string{
		"message": fmt.Sprintf("game created with id: %v created with ease", createdGameId),
	}
	b, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
