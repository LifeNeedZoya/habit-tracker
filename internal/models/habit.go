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
	Color       string `json:"color" gorm:"default:#3b82f6"`

	FrequencyType     FrequencyType `json:"frequency_type" gorm:"not null" validate:"required"`
	FrequencyValue    int           `json:"frequency_value" gorm:"default:1"`
	FrequencyInterval string        `json:"frequency_interval" gorm:"default:day"`
	FlexibleMin       int           `json:"flexible_min" gorm:"default:0"`
	FlexibleMax       int           `json:"flexible_max" gorm:"default:0"`

	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	HabitLogs []HabitLog `json:"habit_logs,omitempty" gorm:"foreignKey:HabitID"`
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
