package database

import (
	"log"

	"github.com/lifeneedzoya/habit-tracker/config"
	"github.com/lifeneedzoya/habit-tracker/pkg/database"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.LoadConfig()
	if cfg == nil {
		panic("Failed to load configuration")
	}

	var err error
	DB, err = database.NewDB(cfg)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

}
