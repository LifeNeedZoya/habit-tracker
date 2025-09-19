package repository

import (
	"time"

	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"gorm.io/gorm"
)

type SessionRepository interface {
	DeleteSession(id string) error
	StopSession(id string, stoppedAt time.Duration) error
	// CheckSessionOwner(ownerId, sessionId string) error
	CompleteSession(id string) error
	SkipSession(id string) error
	// GetUserSessions(userId string) ([]models.Session, error)
	// CreateActivity(dto dto.ActivityDTO, userID uint) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func CreateSession(tx *gorm.DB, session models.Session) error {
	return tx.Create(&session).Error
}

func (r *sessionRepository) DeleteSession(id string) error {
	return r.db.Delete(&models.Session{}, "id = ?", id).Error
}

func (r *sessionRepository) StopSession(id string, stoppedAt time.Duration) error {
	var session models.Session
	err := r.db.Where("deleted_at = ?", nil).First(&session, "id = ?", id).Error

	if err != nil {
		return err
	}
	rightNow := time.Now()
	session.PausedAt = &rightNow
	session.PausedTime += 1

	return r.db.Save(&session).Error

}

func (r *sessionRepository) CompleteSession(id string) error {
	var session models.Session

	err := r.db.Take(&session, "id = ?", id).Error
	if err != nil {
		return err
	}

	session.Status = models.SessionStatus(models.SessionCompleted)

	res := r.db.Save(session)
	if res.Error != nil {
		return err
	}

	return nil
}

func (r *sessionRepository) SkipSession(id string) error {
	var session models.Session

	if res := r.db.Take(&session, id); res.Error != nil {
		return nil
	}

	session.Status = models.SessionSkipped

	if err := r.db.Save(session).Error; err != nil {
		return err
	}

	return nil

}
