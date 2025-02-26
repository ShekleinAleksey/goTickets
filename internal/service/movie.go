package service

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/ShekleinAleksey/goTickets/internal/repository"
)

type MovieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetAllMovies() ([]entity.Movie, error) {
	movies, err := s.repo.GetAllMovies()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) GetMovieByID(id int) (entity.Movie, error) {
	movie, err := s.repo.GetMovieByID(id)
	if err != nil {
		return entity.Movie{}, err
	}
	return movie, nil
}
