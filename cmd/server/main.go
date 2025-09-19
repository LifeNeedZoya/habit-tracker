package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/config"
	"github.com/lifeneedzoya/habit-tracker/internal/api"
	"github.com/lifeneedzoya/habit-tracker/internal/handlers"
	"github.com/lifeneedzoya/habit-tracker/internal/repository"
	"github.com/lifeneedzoya/habit-tracker/internal/services"
	"github.com/lifeneedzoya/habit-tracker/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := database.ConnectDB(dsn)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	habitRepo := repository.NewHabitRepository(db)
	habitService := services.NewHabitService(habitRepo)
	habitHandler := handlers.NewHabitHandler(&habitService)

	sessionRepo := repository.NewSessionRepository(db)
	sessionService := services.NewSessionService(sessionRepo)
	sessionHandler := handlers.NewSessionHandler(sessionService)

	activityRepo := repository.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepo)
	activityHandler := handlers.NewActivityHandler(activityService)

	r := gin.Default()
	api.SetupRoutes(r, userHandler, habitHandler, sessionHandler, activityHandler)
	r.Run(cfg.ServerAddress)
}
