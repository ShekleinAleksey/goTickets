package entity

type Movie struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Genre       string `json:"genre"`
}
