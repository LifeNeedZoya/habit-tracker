package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/handlers"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	v1 := r.Group("/api/v1")
	{
		userRoutes := v1.Group("/users")
		{

			userRoutes.GET("/", userHandler.GetAllUsers)
			userRoutes.POST("/", userHandler.CreateUser)
			userRoutes.POST("/login", userHandler.LoginUser)
		}

		// habitRoutes := userRoutes.Group("/:userId/habits")
		// habitRoutes.Use()
		// {
		// 	habitRoutes.GET("/", userHandler.GetAllHabits)
		// 	habitRoutes.POST("/", userHandler.CreateHabit)
		// 	habitRoutes.PUT("/:habitId", userHandler.UpdateHabit)
		// 	habitRoutes.DELETE("/:habitId", userHandler.DeleteHabit)
		// }
	}
}
