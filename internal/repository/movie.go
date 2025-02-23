package repository

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/jmoiron/sqlx"
)

type MovieRepository struct {
	db *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetAllMovies() ([]entity.Movie, error) {
	var movies []entity.Movie

	err := r.db.Select(&movies, "SELECT * FROM movies")
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepository) GetMovieByID(id int) (entity.Movie, error) {
	var movie entity.Movie

	err := r.db.Get(&movie, "SELECT * FROM movies WHERE id = $1", id)
	if err != nil {
		return entity.Movie{}, err
	}

	return movie, nil
}
