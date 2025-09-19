package dto

import (
	"time"

	"github.com/lifeneedzoya/habit-tracker/internal/models"
)

type FrequencyType string

const (
	FrequencyDaily    FrequencyType = "daily"
	FrequencyWeekly   FrequencyType = "weekly"
	FrequencyMonthly  FrequencyType = "monthly"
	FrequencyCustom   FrequencyType = "custom"
	FrequencyFlexible FrequencyType = "flexible"
)

type CreateHabitDTO struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"omitempty"`

	FrequencyType   FrequencyType `json:"frequency_type" binding:"required"`
	FrequencyValue  int           `json:"frequency_value" binding:"required"`
	FrequencyInWeek int           `json:"frequency_in_week" binding:"required,max=7" `
}

type HabitDTO struct {
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"omitempty"`

	FrequencyType   FrequencyType `json:"frequency_type" binding:"required"`
	FrequencyValue  int           `json:"frequency_value" binding:"required"`
	FrequencyInWeek int           `json:"frequency_in_week" binding:"required,max=7" `

	IsActive  bool      `json:"is_active" `
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateHabitDTO struct {
	Name        string `json:"name" binding:"omitempty,min=3"`
	Description string `json:"description" binding:"omitempty"`

	FrequencyType   FrequencyType `json:"frequency_type" binding:"omitempty"`
	FrequencyValue  *int          `json:"frequency_value" binding:"omitempty"`
	FrequencyInWeek *int          `json:"frequency_in_week" binding:"omitempty,max=7" `
}

func ToHabitResponse(habit *models.Habit) HabitDTO {
	return HabitDTO{
		ID:              habit.ID,
		Name:            habit.Name,
		Description:     habit.Description,
		CreatedAt:       habit.CreatedAt,
		FrequencyType:   FrequencyType(habit.FrequencyType),
		FrequencyValue:  habit.FrequencyValue,
		FrequencyInWeek: habit.FrequencyInWeek,
	}
}
