package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ShekleinAleksey/goTickets/internal/model"
)

var movieScreenings []model.MovieScreening

func CreateMovieScreening(w http.ResponseWriter, r *http.Request) {
	var screening model.MovieScreening
	if err := json.NewDecoder(r.Body).Decode(&screening); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	screening.ID = len(movieScreenings) + 1
	movieScreenings = append(movieScreenings, screening)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(screening)
}
