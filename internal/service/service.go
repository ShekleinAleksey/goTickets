package service

import "github.com/ShekleinAleksey/goTickets/internal/repository"

type Deps struct {
	UserRepository           *repository.UserRepository
	MovieRepository          *repository.MovieRepository
	MovieScreeningRepository *repository.MovieScreeningRepository
	TicketRepository         *repository.TicketRepository
}

type Service struct {
	UserService           *UserService
	MovieService          *MovieService
	MovieScreeningService *MovieScreeningService
	TicketService         *TicketService
}

func NewService(deps Deps) *Service {
	return &Service{
		UserService:           NewUserService(*deps.UserRepository),
		MovieService:          NewMovieService(*deps.MovieRepository),
		MovieScreeningService: NewMovieScreeningService(*deps.MovieScreeningRepository),
		TicketService:         NewTicketService(*deps.TicketRepository),
	}
}
