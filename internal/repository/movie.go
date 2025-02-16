package repository

import (
	"database/sql"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetAllMovies() ([]entity.Movie, error) {
	var movies []entity.Movie
	var err error
	//Получение всех фильмов
	return movies, err
}
