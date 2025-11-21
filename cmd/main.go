package main

import (
	"log"
	"os"
	"saint-seiya-awakening/internal/config"
	"saint-seiya-awakening/internal/database"
	"saint-seiya-awakening/internal/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.ConnectDb()

	if config.Cfg.RunMigrations == "true" {
		database.MigrateDB()
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.SetupRoutes(router)

	log.Fatal(router.Run(":" + config.Cfg.Port))
}
