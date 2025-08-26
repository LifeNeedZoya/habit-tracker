package main

import (
	"fmt"
	"log"

	"github.com/lifeneedzoya/habit-tracker/config"
	"github.com/lifeneedzoya/habit-tracker/internal/models"
	"github.com/lifeneedzoya/habit-tracker/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := database.ConnectDB(dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Habit{},
		&models.Session{},
		&models.HabitLog{},
		&models.Activity{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully!")
}
