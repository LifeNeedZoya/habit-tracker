package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/dto"
	"github.com/lifeneedzoya/habit-tracker/internal/services"
)

type SessionHandler struct {
	service services.SessionService
}

func NewSessionHandler(service services.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) DeleteSession(c *gin.Context) {
	var session dto.SessionId

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.DeleteSession(session.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func (h *SessionHandler) StopSession(c *gin.Context) {
	var sessionDto dto.StopSession

	if err := c.ShouldBindJSON(&sessionDto); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.StopSession(sessionDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func (h *SessionHandler) CompleteSession(c *gin.Context) {
	var sessionDto dto.SessionId

	if err := c.ShouldBindJSON(&sessionDto); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.CompleteSession(sessionDto.ID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func (h *SessionHandler) SkipSession(c *gin.Context) {
	var sessionDto dto.SessionId

	if err := c.ShouldBindJSON(&sessionDto); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Bad request",
		})
		return
	}

	if err := h.service.SkipSession(sessionDto.ID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
