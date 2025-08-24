package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"github.com/lifeneedzoya/habit-tracker/internal/repository"
	"github.com/lifeneedzoya/habit-tracker/internal/services"
	"gorm.io/gorm" // Assuming GORM is used for the repository
)

type Handler struct {
	UserService *services.UserService
}

// NewHandler creates a new Handler with dependencies
func NewHandler(db *gorm.DB) *Handler {
	repo := &repository.UserRepository{DB: db}
	userService := &services.UserService{Repo: repo}
	return &Handler{UserService: userService}
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers() // Call method on instance
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.UserService.CreateUser(user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
}
