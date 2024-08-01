package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"errors"

	"fmt"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Add(user model.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := u.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return errors.New("error inserting user: %w")
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	query := "SELECT id FROM users WHERE username = $1 AND password = $2"
	row := u.db.QueryRow(query, user.Username, user.Password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return errors.New("error checking user availability: %w")
	}
	return nil // TODO: replace this
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by id: %w", err)
	}

	return &user, nil
}
