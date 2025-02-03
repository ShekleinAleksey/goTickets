package main

import (
	"log"
	"net/http"

	"github.com/ShekleinAleksey/goTickets/internal/handler"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	//movie
	router.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	router.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	//user
	router.HandleFunc("/screenings", handler.CreateUser).Methods("POST")
	//movie_screening
	router.HandleFunc("/screenings", handler.CreateMovieScreening).Methods("POST")
	//ticket
	router.HandleFunc("/tickets/get", handler.GetTicket).Methods("GET")
	router.HandleFunc("/tickets/buy", handler.BuyTicket).Methods("POST")
	router.HandleFunc("/tickets/user/{id}", handler.GetUserTickets).Methods("GET")

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
