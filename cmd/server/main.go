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

	r := gin.Default()
	api.SetupRoutes(r, userHandler)
	r.Run(cfg.ServerAddress)
}
