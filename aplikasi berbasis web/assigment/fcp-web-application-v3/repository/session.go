package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	filebasedDb *filebased.Data
}

func NewSessionsRepo(filebasedDb *filebased.Data) *sessionsRepo {
	return &sessionsRepo{filebasedDb}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	// Load existing sessions
	sessions, err := u.loadSessions()
	if err != nil {
		return err
	}

	// Append new session
	sessions = append(sessions, session)

	// Save sessions back to file
	return u.saveSessions(sessions)
}

func (u *sessionsRepo) DeleteSession(token string) error {
	// Load existing sessions
	sessions, err := u.loadSessions()
	if err != nil {
		return err
	}

	// Filter out the session with the given token
	filteredSessions := []model.Session{}
	for _, s := range sessions {
		if s.Token != token {
			filteredSessions = append(filteredSessions, s)
		}
	}

	// Save the filtered sessions back to file
	return u.saveSessions(filteredSessions)
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	// Load existing sessions
	sessions, err := u.loadSessions()
	if err != nil {
		return err
	}

	// Update the session with the same email
	updated := false
	for i, s := range sessions {
		if s.Email == session.Email {
			sessions[i] = session
			updated = true
			break
		}
	}

	if !updated {
		return errors.New("session not found")
	}

	// Save updated sessions back to file
	return u.saveSessions(sessions)
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	// Load existing sessions
	sessions, err := u.loadSessions()
	if err != nil {
		return model.Session{}, err
	}

	// Search for session with the given email
	for _, s := range sessions {
		if s.Email == email {
			return s, nil
		}
	}

	return model.Session{}, errors.New("session not found")
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	// Load existing sessions
	sessions, err := u.loadSessions()
	if err != nil {
		return model.Session{}, err
	}

	// Search for session with the given token
	for _, s := range sessions {
		if s.Token == token {
			return s, nil
		}
	}

	return model.Session{}, errors.New("session not found")
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}

// Helper function to load sessions from file
func (u *sessionsRepo) loadSessions() ([]model.Session, error) {
	file, err := os.Open(u.filebasedDb.DB.Path())
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Session{}, nil // If file does not exist, return an empty slice
		}
		return nil, err
	}
	defer file.Close()

	var sessions []model.Session
	err = json.NewDecoder(file).Decode(&sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// Helper function to save sessions to file
func (u *sessionsRepo) saveSessions(sessions []model.Session) error {
	file, err := os.Create(u.filebasedDb.DB.Path())
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(&sessions)
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, errors.New("token expired")
	}

	return session, nil
}
