package services

import (
	"fmt"

	"github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"github.com/lifeneedzoya/habit-tracker/internal/repository"
)

type HabitService interface {
	CreateHabit(habit dto.CreateHabitDTO, userId uint) error
	GetAllHabits(userId uint) ([]models.Habit, error)
	GetHabitById(id string) (dto.HabitDTO, error)
	UpdateHabit(id string, dto dto.UpdateHabitDTO) error
	DeleteHabit(id string) error
}

type habitService struct {
	repo repository.HabitRepository
}

func NewHabitService(repo repository.HabitRepository) HabitService {
	return &habitService{repo: repo}
}

func (s *habitService) CreateHabit(habit dto.CreateHabitDTO, userId uint) error {

	return s.repo.CreateHabit(&habit, userId)
}

func (s *habitService) GetAllHabits(userId uint) ([]models.Habit, error) {
	habits, err := s.repo.GetAllHabits(userId)

	if err != nil {
		return nil, err
	}

	return habits, nil
}

func (s *habitService) GetHabitById(id string) (dto.HabitDTO, error) {
	var habitDTO dto.HabitDTO
	habit, err := s.repo.GetHabitById(id)

	if err != nil {
		return habitDTO, fmt.Errorf("habit does not exist")
	}

	habitDTO = dto.ToHabitResponse(habit)
	return habitDTO, nil

}

func (s *habitService) UpdateHabit(id string, dto dto.UpdateHabitDTO) error {
	err := s.repo.UpdateHabit(id, dto)

	if err != nil {
		return err
	}

	return nil
}

func (s *habitService) DeleteHabit(id string) error {

	err := s.repo.DeleteHabit(id)

	if err != nil {
		return err
	}

	return nil
}
