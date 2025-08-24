package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/api"
	"github.com/lifeneedzoya/habit-tracker/config"
	"github.com/lifeneedzoya/habit-tracker/pkg/database"
)

func main() {
	cfg := config.LoadConfig()
	if cfg == nil {
		panic("Failed to load configuration")
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	api.SetupRoutes(router, db)
	router.Run(cfg.ServerAddress)
}

// func main() {
// 	cfg := config.LoadConfig()
// 	router := gin.Default()
// 	api.SetupRoutes(router, db)
// 	router.Run(cfg.ServerAddress)
// }
