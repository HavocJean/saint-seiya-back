package main

import (
	"log"
	"saint-seiya-awakening/internal/config"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.ConnectDb()

	if config.Cfg.RunMigrations == "true" {
		database.MigrateDB()
	}

	router := gin.Default()
	routes.SetupRoutes(router)

	log.Fatal(router.Run(":" + config.Cfg.Port))
}
