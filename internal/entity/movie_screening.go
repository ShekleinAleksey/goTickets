package entity

import "time"

//Сеанс показа фильма
type MovieScreening struct {
	ID       int
	Movie    Movie
	ShowTime time.Time
	Hall     int
}
