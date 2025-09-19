package models

import (
	"time"

	"gorm.io/gorm"
)

type FrequencyType string

const (
	FrequencyDaily    FrequencyType = "daily"
	FrequencyWeekly   FrequencyType = "weekly"
	FrequencyMonthly  FrequencyType = "monthly"
	FrequencyCustom   FrequencyType = "custom"
	FrequencyFlexible FrequencyType = "flexible"
)

type Habit struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null" validate:"required,min=1,max=100"`
	Description string `json:"description" gorm:"type:text"`

	FrequencyType   FrequencyType `json:"frequency_type" gorm:"not null" validate:"required"`
	FrequencyValue  int           `json:"frequency_value" gorm:"default:1"`
	FrequencyInWeek int           `json:"frequency_in_week" `

	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	HabitLogs []HabitLog `json:"habit_logs,omitempty" gorm:"foreignKey:HabitID"`
	UserID    uint       `json:"user_id" gorm:"not null"`
	User      User       `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type HabitLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	HabitID   uint      `json:"habit_id" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"not null"`
	Completed bool      `json:"completed" gorm:"default:false"`
	Notes     string    `json:"notes" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Habit Habit `json:"habit,omitempty" gorm:"foreignKey:HabitID"`
}
