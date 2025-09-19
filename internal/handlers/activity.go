package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/services"
	"github.com/lifeneedzoya/habit-tracker/internal/utils"
)

type ActivityHandler struct {
	service services.ActivityService
}

func NewActivityHandler(service services.ActivityService) *ActivityHandler {
	return &ActivityHandler{service: service}
}

func (h *ActivityHandler) CreateActivities(c *gin.Context) {
	var activityDTO dto.ActivityDTO

	if err := c.ShouldBindJSON(&activityDTO); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := utils.GetUserIDFromContext(c)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Authentication error",
		})
		return
	}

	userIDAsInt, err := strconv.Atoi(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}

	if err := h.service.CreateActivity(activityDTO, userIDAsInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func (h *ActivityHandler) GetUserActivities(c *gin.Context) {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Authentication error",
		})
		return
	}

	activities, err := h.service.GetUserActivities(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    activities,
	})

}

func (h *ActivityHandler) CompleteActivity(c *gin.Context) {
	var id string

	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured"})
		return
	}

	err = h.service.CompleteActivity(idAsInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "completed",
	})

}
