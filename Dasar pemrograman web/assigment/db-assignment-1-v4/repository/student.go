package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"errors"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
}

type studentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	return []model.Student{}, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	row := s.db.QueryRow("SELECT id, name, address, class FROM students WHERE id = $1", id)

	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	query := "INSERT INTO students (id, name, address, class) VALUES (?, ?, ?, ?)"
    _, err := s.db.Exec(query, student.ID, student.Name, student.Address, student.Class)
    if err != nil {
        return err
    }
    return nil // TODO: replace this 
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	query := "UPDATE students SET name = ?, address = ?, class = ? WHERE id = ?"
    result, err := s.db.Exec(query, student.Name, student.Address, student.Class, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return errors.New("no student found with the given id")
    }

    return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	query := "DELETE FROM students WHERE id = ?"
    _, err := s.db.Exec(query, id)
    if err != nil {
        return err
    }
    return nil // TODO: replace this
}
