package main

import (
	"os"
	"saint-seiya-back/internal/bootstrap"
	"saint-seiya-back/internal/config"
	"saint-seiya-back/internal/infrastructure/database"
	"saint-seiya-back/internal/infrastructure/http/middleware"
	"saint-seiya-back/internal/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	config.Load()
	log.Info().Msg("Starting application...")

	database.ConnectDb()
	log.Info().Msg("Database connected")

	if config.Cfg.RunMigrations == "true" {
		database.MigrateDB()
		log.Info().Msg("Database migrations completed")
	}

	app := bootstrap.InitApp()

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.SetupRoutes(router, app)

	log.Info().Str("port", config.Cfg.Port).Msg("Server starting")
	router.Run(":" + config.Cfg.Port)
}
