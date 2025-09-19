package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/services"
	"github.com/lifeneedzoya/habit-tracker/internal/utils"
)

type HabitHandler struct {
	service services.HabitService
}

func NewHabitHandler(service *services.HabitService) *HabitHandler {
	return &HabitHandler{service: *service}
}

func (h *HabitHandler) CreateHabit(c *gin.Context) {
	var habit dto.CreateHabitDTO

	userID, exists := utils.GetStrValFromContext(c, "user_id")

	if !exists {
		log.Fatal("user_id does not exists")
	}

	userIDAsInt, err := strconv.Atoi(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}

	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.service.CreateHabit(habit, uint(userIDAsInt)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created successfully",
	})

}
func (h *HabitHandler) GetAllHabits(c *gin.Context) {

	userID, exists := utils.GetStrValFromContext(c, "user_id")

	if !exists {
		log.Fatal("user_id does not exists")
	}

	userIDAsInt, err := strconv.Atoi(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}

	habits, err := h.service.GetAllHabits(uint(userIDAsInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get user habits successfully",
		"data":    habits,
	})

}

func (h *HabitHandler) GetHabitById(c *gin.Context) {
	var habitId string
	if err := c.ShouldBindJSON(&habitId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	habitDTO, err := h.service.GetHabitById(habitId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": habitDTO})

}

func (h *HabitHandler) DeleteHabit(c *gin.Context) {
	habitId := c.Param("habitId")

	err := h.service.DeleteHabit(habitId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})

}

func (h *HabitHandler) UpdateHabit(c *gin.Context) {
	var habitDTO dto.UpdateHabitDTO
	habitId := c.Param("habitId")

	if err := c.ShouldBindJSON(&habitDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid information"})
		return
	}

	err := h.service.UpdateHabit(habitId, habitDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})

}
