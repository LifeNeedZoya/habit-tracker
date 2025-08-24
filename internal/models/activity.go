package models

import (
	"time"

	"gorm.io/gorm"
)

type ActivityType string

const (
	ActivityPlanned   ActivityType = "planned"
	ActivityUnplanned ActivityType = "unplanned"
)

type Activity struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"not null" validate:"required,min=1,max=100"`
	Description string       `json:"description" gorm:"type:text"`
	Type        ActivityType `json:"type" gorm:"not null" validate:"required"`
	Color       string       `json:"color" gorm:"default:#10b981"`

	Duration  int        `json:"duration"` // in minutes
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Date      time.Time  `json:"date" gorm:"not null"`

	SessionCount      int `json:"session_count" gorm:"default:1"`        // number of sessions
	SessionDuration   int `json:"session_duration" gorm:"default:25"`    // minutes per session
	BreakDuration     int `json:"break_duration" gorm:"default:5"`       // minutes between sessions
	LongBreakAfter    int `json:"long_break_after" gorm:"default:4"`     // long break after N sessions
	LongBreakDuration int `json:"long_break_duration" gorm:"default:15"` // long break duration

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Sessions []Session `json:"sessions,omitempty" gorm:"foreignKey:ActivityID"`
}
