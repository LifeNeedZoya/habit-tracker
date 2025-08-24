package repository

import (
	"github.com/lifeneedzoya/habit-tracker/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) Create(user models.User) error {
	return r.DB.Create(&user).Error
}
