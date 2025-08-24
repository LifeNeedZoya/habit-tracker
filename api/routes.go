package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	handler := handlers.NewHandler(db)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Habit Tracker API"})
		})
		v1.POST("/users", handler.CreateUser)
	}
}
