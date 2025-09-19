package repository

import (
	"log"

	dto "github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"gorm.io/gorm"
)

type HabitRepository interface {
	CreateHabit(habit *dto.CreateHabitDTO, userId uint) error
	GetAllHabits(userId uint) ([]models.Habit, error)
	GetHabitById(id string) (*models.Habit, error)
	UpdateHabit(id string, dto dto.UpdateHabitDTO) error
	DeleteHabit(id string) error
}

type habitRepository struct {
	Db *gorm.DB
}

func NewHabitRepository(db *gorm.DB) HabitRepository {
	return &habitRepository{Db: db}
}

func (r *habitRepository) CreateHabit(dto *dto.CreateHabitDTO, userId uint) error {
	habit := models.Habit{
		Name:            dto.Name,
		Description:     dto.Description,
		FrequencyType:   models.FrequencyType(dto.FrequencyType),
		FrequencyValue:  dto.FrequencyValue,
		FrequencyInWeek: dto.FrequencyInWeek,
		UserID:          userId,
	}
	err := r.Db.Create(&habit).Error

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (r *habitRepository) GetAllHabits(userId uint) ([]models.Habit, error) {
	var habits []models.Habit
	err := r.Db.Where("deleted_at = ?", nil).Find(&habits, "user_id = ?", userId).Error
	return habits, err
}

func (r *habitRepository) GetHabitById(id string) (*models.Habit, error) {
	var habit models.Habit
	err := r.Db.Where("deleted_at = ?", nil).First(&habit, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &habit, nil
}

func (r *habitRepository) UpdateHabit(id string, dto dto.UpdateHabitDTO) error {

	habit, err := r.GetHabitById(id)
	if err != nil {
		return err
	}

	if dto.Name != "" {
		habit.Name = dto.Name
	}

	if dto.Description != "" {
		habit.Description = dto.Description
	}

	if dto.FrequencyInWeek != nil {
		habit.FrequencyInWeek = *dto.FrequencyInWeek

	}
	if dto.FrequencyValue != nil {
		habit.FrequencyValue = *dto.FrequencyValue

	}
	return r.Db.Save(habit).Error
}

func (r *habitRepository) DeleteHabit(id string) error {

	return r.Db.Delete(&models.Habit{}, "id = ?", id).Error
}
