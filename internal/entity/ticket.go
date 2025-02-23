package entity

type Ticket struct {
	ID             int            `json:"id" db:"id"`
	MovieScreening MovieScreening `json:"movie_screening" db:"movie_screening"`
	UserID         int            `json:"user_id" db:"user_id"`
	Price          float64        `json:"price" db:"price"`
	Seat           int            `json:"seat" db:"seat"`
}
