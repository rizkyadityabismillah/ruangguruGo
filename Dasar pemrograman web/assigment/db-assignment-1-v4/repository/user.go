package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
    "errors"
	
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
	query := "INSERT INTO users (id, username, password) VALUES (?, ?, ?)"
    _, err := u.db.Exec(query, user.ID, user.Username, user.Password)
    if err != nil {
        return err
    }
    return nil
	// TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	query := "SELECT id, username, password FROM users WHERE username = ? AND password = ?"
    var existingUser model.User
    row := u.db.QueryRow(query, user.Username, user.Password)
    err := row.Scan(&existingUser.ID, &existingUser.Username, &existingUser.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return errors.New("user not found")
        }
        return err
    }
    return nil
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
