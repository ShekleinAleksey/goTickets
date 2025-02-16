package repository

import (
	"database/sql"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entity.User) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id", user.Name, user.Balance).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
