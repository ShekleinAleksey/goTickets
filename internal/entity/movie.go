package entity

type Movie struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Duration    int    `json:"duration" db:"duration"`
	Genre       string `json:"genre" db:"genre"`
}
