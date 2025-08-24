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

type SessionType string

const (
	SessionWork      SessionType = "work"
	SessionBreak     SessionType = "break"
	SessionLongBreak SessionType = "long_break"
)

type Session struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	ActivityID uint `json:"activity_id" gorm:"not null"`

	Type            SessionType   `json:"type" gorm:"not null"`
	Status          SessionStatus `json:"status" gorm:"default:pending"`
	PlannedDuration int           `json:"planned_duration"`
	ActualDuration  int           `json:"actual_duration"`

	StartTime  *time.Time `json:"start_time"`
	EndTime    *time.Time `json:"end_time"`
	PausedAt   *time.Time `json:"paused_at"`
	PausedTime int        `json:"paused_time"`

	SessionNumber  int  `json:"session_number" gorm:"default:1"`
	IsBreakSkipped bool `json:"is_break_skipped" gorm:"default:false"`

	Notes     string         `json:"notes" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Activity Activity `json:"activity,omitempty" gorm:"foreignKey:ActivityID"`
}
