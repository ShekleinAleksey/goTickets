package entity

import "time"

//Сеанс показа фильма
type MovieScreening struct {
	ID             int       `json:"id"`
	Movie          Movie     `json:"movie"`
	ShowTime       time.Time `json:"show_time"`
	Hall           int       `json:"hall"`
	AvailableSeats int       `json:"available_seats"`
}
