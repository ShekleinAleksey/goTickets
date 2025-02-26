package entity

import "time"

// Сеанс показа фильма
type MovieScreening struct {
	ID             int       `json:"id" db:"id"`
	MovieID        int       `json:"movie_id" db:"movie_id"`
	Movie          Movie     `json:"movie" db:"movie"`
	ShowTime       time.Time `json:"show_time" db:"show_time"`
	Hall           int       `json:"hall" db:"hall"`
	TicketPrice    float64   `json:"ticket_price" db:"ticket_price"`
	AvailableSeats int       `json:"available_seats" db:"available_seats"`
}
