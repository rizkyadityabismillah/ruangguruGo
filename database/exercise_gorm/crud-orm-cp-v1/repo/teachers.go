package repo

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

// Save stores a single Teacher record in the database
func (t TeacherRepo) Save(data model.Teacher) error {
	if err := t.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

// Query retrieves all Teacher records from the database
func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teachers []model.Teacher
	if err := t.db.Select("*").Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

// Update modifies the name of a Teacher record with the given ID
func (t TeacherRepo) Update(id uint, name string) error {
	if err := t.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

// Delete removes the Teacher record with the given ID from the database
func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	if err := t.db.Where("id = ?", id).Delete(&teacher);err.Error != nil {
		return fmt.Errorf("ERROR DELETE Teacher")
	}
	return nil
}
