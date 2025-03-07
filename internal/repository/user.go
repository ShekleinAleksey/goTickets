package repository

import (
	"github.com/ShekleinAleksey/goTickets/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	var createdUser entity.User
	err := r.db.QueryRow("INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id, name, balance", user.Name, user.Balance).Scan(&createdUser.ID, &createdUser.Name, &createdUser.Balance)
	if err != nil {
		return entity.User{}, err
	}
	return createdUser, nil
}

func (r *UserRepository) GetUserByID(id int) (entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *entity.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, balance = $2 WHERE id = $3", user.Name, user.Balance, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
