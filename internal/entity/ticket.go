package entity

type Ticket struct {
	ID             int            `json:"id"`
	MovieScreening MovieScreening `json:"movie_screening"`
	UserID         int            `json:"user_id"`
	Price          float64        `json:"price"`
	Seat           string         `json:"seat"`
}
