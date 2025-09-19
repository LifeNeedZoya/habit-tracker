package services

import (
	"errors"
	"fmt"

	"github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"github.com/lifeneedzoya/habit-tracker/internal/repository"
)

type ActivityService interface {
	CreateActivity(DTO dto.ActivityDTO, userID int) error
	GetUserActivities(userId string) ([]models.Activity, error)
	CheckActivityOwner(ownerId, sessionId int) error
	CompleteActivity(activityID int) error
}

type activityService struct {
	repo repository.ActivityRepository
}

func NewActivityService(repo repository.ActivityRepository) ActivityService {
	return &activityService{repo: repo}
}

func (s *activityService) CreateActivity(DTO dto.ActivityDTO, userID int) error {

	if err := s.repo.CreateActivity(DTO, uint(userID)); err != nil {
		return errors.New("error occured during creating session")
	}

	return nil
}

func (s *activityService) GetUserActivities(userId string) ([]models.Activity, error) {
	activities, err := s.repo.GetUserActivities(userId)

	if err != nil {
		return []models.Activity{}, err
	}

	return activities, nil
}

func (s *activityService) CheckActivityOwner(ownerId, sessionId int) error {

	if err := s.repo.CheckActivityOwner(ownerId, sessionId); err != nil {
		return fmt.Errorf("failed to verify session ownership: %w", err)
	}

	return nil
}

func (s *activityService) CompleteActivity(activityID int) error {

	if err := s.repo.CompleteActivity(activityID); err != nil {
		return fmt.Errorf("failed to verify session ownership: %w", err)
	}

	return nil
}
