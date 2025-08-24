package database

import (
	"fmt"

	"github.com/lifeneedzoya/habit-tracker/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
