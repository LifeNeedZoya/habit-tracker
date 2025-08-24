package services

import (
	"errors"

	"github.com/lifeneedzoya/habit-tracker/internal/repository"

	"github.com/lifeneedzoya/habit-tracker/internal/models"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) CreateUser(user models.User) error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	return s.Repo.Create(user)
}
