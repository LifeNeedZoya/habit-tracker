package services

import (
	"errors"

	"github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/repository"
)

type SessionService interface {
	DeleteSession(id string) error
	StopSession(dto dto.StopSession) error
	CompleteSession(id string) error
	SkipSession(id string) error
}

type sessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) SessionService {
	return &sessionService{repo: repo}
}

func (s *sessionService) DeleteSession(id string) error {
	err := s.repo.DeleteSession(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionService) StopSession(dto dto.StopSession) error {
	if err := s.repo.StopSession(dto.ID, dto.StoppedAt); err != nil {
		return errors.New("error occured during stopping session")
	}

	return nil
}

func (s *sessionService) CompleteSession(id string) error {
	if err := s.repo.CompleteSession(id); err != nil {
		return errors.New("error occured during completing session")
	}
	return nil
}

func (s *sessionService) SkipSession(id string) error {
	if err := s.repo.SkipSession(id); err != nil {
		return errors.New("error occured during skipping session")
	}
	return nil
}
