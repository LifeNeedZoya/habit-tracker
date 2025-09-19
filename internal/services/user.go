package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"github.com/lifeneedzoya/habit-tracker/internal/repository"
	"github.com/lifeneedzoya/habit-tracker/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user dto.CreateUserRequest) error
	AuthenticateUser(email, password string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) CreateUser(user dto.CreateUserRequest) error {
	fmt.Println(user)

	if user.Name == "" {
		return errors.New("name cannot be empty")
	}

	var oldUser models.User

	err := s.repo.FindByEmail(&oldUser, user.Email)

	if err == nil {
		return errors.New("user already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPass)

	fmt.Println("hashedPass : " + string(hashedPass))
	fmt.Println("user.pass : " + user.Password)

	return s.repo.CreateUser(
		&models.User{Name: user.Name,
			Password: user.Password,
			Email:    user.Email},
	)
}

func (s *userService) AuthenticateUser(email, password string) (string, error) {
	var user models.User

	if err := s.repo.FindByEmail(&user, email); err != nil {
		log.Printf("User lookup failed: %v", err)
		return "", errors.New("authentication failed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("authentication failed")
	}

	token, err := utils.CreateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		log.Printf("Token creation failed: %v", err)
		return "", errors.New("authentication failed")
	}

	return token, nil
}
