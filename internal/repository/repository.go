package repository

import "database/sql"

type Repository struct {
	UserRepository           *UserRepository
	MovieRepository          *MovieRepository
	MovieScreeningRepository *MovieScreeningRepository
	TicketRepository         *TicketRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository:           NewUserRepository(db),
		MovieRepository:          NewMovieRepository(db),
		MovieScreeningRepository: NewMovieScreeningRepository(db),
		TicketRepository:         NewTicketRepository(db),
	}
}
