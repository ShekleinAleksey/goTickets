package service

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/ShekleinAleksey/goTickets/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *entity.User) (int, error) {
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserService) GetUserByID(id int) (entity.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(user *entity.User) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
