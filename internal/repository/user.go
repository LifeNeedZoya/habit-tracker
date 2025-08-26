package repository

import (
	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	CreateUser(user *models.User) error
	FindByEmail(user *models.User, email string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindByEmail(user *models.User, email string) error {
	return r.db.First(user, "email = ?", email).Error
}
