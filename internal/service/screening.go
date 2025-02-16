package service

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/ShekleinAleksey/goTickets/internal/repository"
)

type MovieScreeningService struct {
	repo repository.MovieScreeningRepository
}

func NewMovieScreeningService(repo repository.MovieScreeningRepository) *MovieScreeningService {
	return &MovieScreeningService{repo: repo}
}

func (s *MovieScreeningService) GetAllMovieScreenings() ([]entity.MovieScreening, error) {
	return s.repo.GetAllMovieScreenings()
}
