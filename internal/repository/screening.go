package repository

import (
	"database/sql"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
)

type MovieScreeningRepository struct {
	db *sql.DB
}

func NewMovieScreeningRepository(db *sql.DB) *MovieScreeningRepository {
	return &MovieScreeningRepository{db: db}
}

func (r *MovieScreeningRepository) GetAllMovieScreenings() ([]entity.MovieScreening, error) {
	var screenings []entity.MovieScreening
	var err error
	//Получение всех сеансов
	return screenings, err
}
