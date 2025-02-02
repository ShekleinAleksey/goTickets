package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ShekleinAleksey/goTickets/internal/model"
)

var movies []model.Movie

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	movie.ID = len(movies) + 1
	movies = append(movies, movie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
