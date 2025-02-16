package repository

import (
	"database/sql"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) GetTicket(id int) (entity.Ticket, error) {
	var ticket entity.Ticket
	var err error
	//Получение билета по id
	return ticket, err
}
