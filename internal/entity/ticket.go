package entity

type Ticket struct {
	ID             int
	MovieScreening MovieScreening
	UserID         int
	Price          float64
	Seat           string
}
