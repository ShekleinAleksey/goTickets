package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository           *UserRepository
	MovieRepository          *MovieRepository
	MovieScreeningRepository *MovieScreeningRepository
	TicketRepository         *TicketRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:           NewUserRepository(db),
		MovieRepository:          NewMovieRepository(db),
		MovieScreeningRepository: NewMovieScreeningRepository(db),
		TicketRepository:         NewTicketRepository(db),
	}
}
