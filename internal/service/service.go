package service

import "github.com/ShekleinAleksey/goTickets/internal/repository"

type Service struct {
	UserService           *UserService
	MovieService          *MovieService
	MovieScreeningService *MovieScreeningService
	TicketService         *TicketService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:           NewUserService(*repo.UserRepository),
		MovieService:          NewMovieService(*repo.MovieRepository),
		MovieScreeningService: NewMovieScreeningService(*repo.MovieScreeningRepository),
		TicketService:         NewTicketService(*repo.TicketRepository),
	}
}
