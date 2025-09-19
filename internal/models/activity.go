package models

import (
	"time"

	"gorm.io/gorm"
)

type ActivityType string

type Activity struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null" validate:"required,min=1,max=100"`

	TotalDuration *time.Duration `json:"duration"`
	StartTime     *time.Time     `json:"start_time"`
	EndTime       *time.Time     `json:"end_time"`

	SessionCount int `json:"session_count" gorm:"default:1"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	UserID   uint `json:"user_id" gorm:"not null"`
	User     User `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sessions []Session
}
