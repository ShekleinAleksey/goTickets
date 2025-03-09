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
	screenings, err := s.repo.GetAllMovieScreenings()
	if err != nil {
		return nil, err
	}
	return screenings, nil
}

func (s *MovieScreeningService) GetScreening(id int) (entity.MovieScreening, error) {
	screening, err := s.repo.GetScreening(id)
	if err != nil {
		return entity.MovieScreening{}, err
	}
	return screening, nil
}

func (s *MovieScreeningService) UpdateScreening(screening *entity.MovieScreening) error {
	err := s.repo.UpdateScreening(screening)
	if err != nil {
		return err
	}
	return nil
}

func (s *MovieScreeningService) FilterScreenings(movieIDStr string, dateStr string) ([]entity.MovieScreening, error) {
	screenings, err := s.repo.FilterScreenings(movieIDStr, dateStr)
	if err != nil {
		return nil, err
	}
	return screenings, nil
}
