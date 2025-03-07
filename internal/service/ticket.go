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

func (s *TicketService) BuyTicket(screeningID int, userID int, seat int) error {

	err := s.repo.CreateTicket(screening, userID, seat)
	if err != nil {
		return err
	}

	// screening, err := h.service.MovieScreeningService.GetScreening(id)
	// err = h.TicketService.BuyTicket(screening, userID, seat)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// user, err := h.service.UserService.GetUserByID(userID)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// user.Balance -= screening.TicketPrice
	// err = h.service.UserService.UpdateUser(&user)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	return nil
}
