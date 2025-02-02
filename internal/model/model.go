package model

import "time"

type Movie struct {
	ID   int
	Name string
}

type User struct {
	ID      int
	Name    string
	Balance float64
}

type Ticket struct {
	ID             int
	MovieScreening MovieScreening
	UserID         int
	Price          float64
	Seat           string
}

//Сеанс показа фильма
type MovieScreening struct {
	ID       int
	Movie    Movie
	ShowTime time.Time
	Hall     int
}
