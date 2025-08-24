package main

import (
	"log"

	"github.com/lifeneedzoya/habit-tracker/internal/database"
	"github.com/lifeneedzoya/habit-tracker/internal/models"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(
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
