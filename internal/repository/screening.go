package repository

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/jmoiron/sqlx"
)

type MovieScreeningRepository struct {
	db *sqlx.DB
}

func NewMovieScreeningRepository(db *sqlx.DB) *MovieScreeningRepository {
	return &MovieScreeningRepository{db: db}
}

func (r *MovieScreeningRepository) GetAllMovieScreenings() ([]entity.MovieScreening, error) {
	var screenings []entity.MovieScreening

	err := r.db.Select(&screenings, "SELECT * FROM movie_screenings")
	if err != nil {
		return nil, err
	}

	return screenings, nil
}

func (r *MovieScreeningRepository) GetScreening(id int) (entity.MovieScreening, error) {
	var screening entity.MovieScreening

	err := r.db.Get(&screening, "SELECT * FROM movie_screenings WHERE id = $1", id)
	if err != nil {
		return entity.MovieScreening{}, err
	}
	return screening, nil
}

func (r *MovieScreeningRepository) UpdateScreening(screening *entity.MovieScreening) error {
	_, err := r.db.Exec("UPDATE movie_screenings SET movie_id = $1, show_time = $2, hall = $3, ticket_price = $4, available_seats = $5 WHERE id = $6", screening.MovieID, screening.ShowTime, screening.Hall, screening.TicketPrice, screening.AvailableSeats, screening.ID)
	if err != nil {
		return err
	}
	return nil
}
