package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/handlers"
	"github.com/lifeneedzoya/habit-tracker/internal/middleware"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, habitHandler *handlers.HabitHandler, sessionHandler *handlers.SessionHandler, activityHandler *handlers.ActivityHandler) {
	v1 := r.Group("/api/v1")
	{
		userRoutes := v1.Group("/users")
		{

			userRoutes.GET("/", userHandler.GetAllUsers)
			userRoutes.POST("/", userHandler.CreateUser)
			userRoutes.POST("/login", userHandler.LoginUser)
		}

		habitRoutes := v1.Group("/habits")
		habitRoutes.Use(middleware.AuthMiddleware())
		{
			habitRoutes.GET("/", habitHandler.GetAllHabits)
			habitRoutes.POST("/", habitHandler.CreateHabit)
			habitRoutes.PUT("/:habitId", habitHandler.GetHabitById)
			habitRoutes.DELETE("/:habitId", habitHandler.DeleteHabit)
		}

		sessionRoutes := v1.Group("/sessions")
		sessionRoutes.Use(middleware.AuthMiddleware())

		{
			sessionRoutes.POST("/complete/:habitId", sessionHandler.CompleteSession)
			sessionRoutes.POST("/stop/:habitId", sessionHandler.StopSession)
			sessionRoutes.POST("/skip/:habitId", sessionHandler.SkipSession)
			sessionRoutes.DELETE("/:habitId", sessionHandler.DeleteSession)
		}

		activityRoutes := v1.Group("/activities")
		activityRoutes.Use(middleware.AuthMiddleware())

		{
			activityRoutes.POST("/", activityHandler.CreateActivities)
			activityRoutes.GET("/", activityHandler.GetUserActivities)
			activityRoutes.POST("/complete", activityHandler.CompleteActivity)
		}

	}
}
