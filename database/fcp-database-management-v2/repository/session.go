package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	result := s.db.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	result := s.db.Where("token = ?", token).Delete(&model.Session{})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	var session model.Session
	result := s.db.Where("username = ?", name).First(&session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("session not found")
	}
	return nil
	  // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	result := s.db.Where("token = ?", token).First(&session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Session{}, result.Error
	}
	return session, nil // TODO: replace this
}
