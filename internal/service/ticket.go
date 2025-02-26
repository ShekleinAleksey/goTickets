package service

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/ShekleinAleksey/goTickets/internal/repository"
)

type TicketService struct {
	repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) GetTicket(id int) (entity.Ticket, error) {
	return s.repo.GetTicket(id)
}

func (s *TicketService) BuyTicket(screening entity.MovieScreening, userID int, seat int) error {

	err := s.repo.CreateTicket(screening, userID, seat)
	if err != nil {
		return err
	}

	return nil
}
