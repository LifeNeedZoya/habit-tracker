package repository

import (
	"errors"
	"time"

	dto "github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	CreateActivity(dto dto.ActivityDTO, userID uint) error
	GetUserActivities(userId string) ([]models.Activity, error)
	CheckActivityOwner(ownerID, activityID int) error
	CompleteActivity(activityID int) error
}

type activityRepository struct {
	DB *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{DB: db}
}

func (r *activityRepository) CreateActivity(dto dto.ActivityDTO, userID uint) error {

	parsedTime, err := time.Parse(time.RFC3339, dto.StartTime)
	if err != nil {
		return err
	}

	totalDuration := time.Duration(dto.TotalDuration)

	tx := r.DB.Begin()

	activity := models.Activity{
		Name:          dto.Name,
		UserID:        userID,
		SessionCount:  len(dto.Sessions),
		StartTime:     &parsedTime,
		CreatedAt:     time.Now(),
		TotalDuration: &totalDuration,
	}
	if err := tx.Create(&activity).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, s := range dto.Sessions {
		session := models.Session{
			Name:       s.Name,
			Order:      s.Order,
			Duration:   s.Duration,
			ActivityID: activity.ID,
			StartTime:  &parsedTime,
		}

		if err := CreateSession(tx, session); err != nil {
			tx.Rollback()
			return err
		}

		parsedTime = parsedTime.Add(time.Duration(s.Duration) * time.Minute)

	}

	return tx.Commit().Error
}

func (r *activityRepository) GetUserActivities(userID string) ([]models.Activity, error) {
	var activities []models.Activity

	if err := r.DB.
		Preload("Sessions").
		Where("end_time IS NULL AND user_id = ?", userID).
		Find(&activities).Error; err != nil {
		return activities, err
	}

	return activities, nil
}

func (r *activityRepository) CheckActivityOwner(ownerID, activityID int) error {
	var activity models.Activity
	if err := r.DB.First(&activity, "user_id = ?", ownerID).Error; err != nil {
		return err
	}

	if activity.UserID != uint(ownerID) {
		return errors.New("access denied")

	}
	return nil
}

func (r *activityRepository) CompleteActivity(activityID int) error {
	var activity models.Activity

	if err := r.DB.First(&activity, "id = ?", activityID).Error; err != nil {
		return err
	}

	RN := time.Now()
	activity.EndTime = &RN

	return nil
}
