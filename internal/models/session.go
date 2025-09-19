package models

import (
	"time"

	"gorm.io/gorm"
)

type SessionStatus string

const (
	SessionPending   SessionStatus = "pending"
	SessionActive    SessionStatus = "active"
	SessionPaused    SessionStatus = "paused"
	SessionCompleted SessionStatus = "completed"
	SessionSkipped   SessionStatus = "skipped"
)

type Session struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Status   SessionStatus `json:"status" gorm:"default:pending"`
	Duration int           `json:"duration"`
	Order    int           `json:"order"`

	StartTime  *time.Time `json:"start_time"`
	EndTime    *time.Time `json:"end_time"`
	PausedAt   *time.Time `json:"paused_at"`
	PausedTime int        `json:"paused_time"`

	Name      string         `json:"name" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	ActivityID uint `json:"activity_id"`
}
