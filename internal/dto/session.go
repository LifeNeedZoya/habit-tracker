package dto

import (
	"time"
)

type SessionStatus string

const (
	SessionPending   SessionStatus = "pending"
	SessionActive    SessionStatus = "active"
	SessionPaused    SessionStatus = "paused"
	SessionCompleted SessionStatus = "completed"
	SessionSkipped   SessionStatus = "skipped"
)

type CreateSessionDTO struct {
	Order    int    `json:"order"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

type sessionDTO struct {
	Order    int    `json:"order"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

type Session struct {
	ID         uint `json:"id"`
	ActivityID uint `json:"activity_id"`
	UserID     uint `json:"user_id"`

	Status          SessionStatus `json:"status"`
	PlannedDuration int           `json:"planned_duration"`
	ActualDuration  int           `json:"actual_duration"`

	StartTime  *time.Time `json:"start_time"`
	EndTime    *time.Time `json:"end_time"`
	PausedAt   *time.Time `json:"paused_at"`
	PausedTime int        `json:"paused_time"`

	SessionNumber int `json:"session_number"`

	CreatedAt time.Time `json:"created_at"`
}

type StopSession struct {
	ID        string        `json:"id"`
	StoppedAt time.Duration `json:"stopped_at"`
}

type SessionId struct {
	ID string `json:"id"`
}
