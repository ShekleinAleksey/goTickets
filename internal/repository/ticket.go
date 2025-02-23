package repository

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TicketRepository struct {
	db *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) GetTicket(id int) (entity.Ticket, error) {
	var ticket entity.Ticket
	err := r.db.Get(&ticket, "SELECT * FROM tickets WHERE id = $1", id)
	if err != nil {
		return entity.Ticket{}, err
	}
	return ticket, nil
}

func (r *TicketRepository) CreateTicket(screening entity.MovieScreening, userID int, seat int) error {

	_, err := r.db.Exec("INSERT INTO tickets (movie_screening_id, user_id, seat, price) VALUES ($1, $2, $3, $4)", screening.ID, userID, seat, screening.TicketPrice)
	if err != nil {
		return err
	}
	return nil
}
