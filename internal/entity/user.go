package entity

type User struct {
	ID      int     `json:"id" db:"id"`
	Name    string  `json:"name" db:"name"`
	Email   string  `json:"email" db:"email"`
	Balance float64 `json:"balance" db:"balance"`
}
