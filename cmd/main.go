package main

import (
	"log"
	"saint-seiya-awakening/internal/config"
	"saint-seiya-awakening/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.ConnectDb()
	
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := config.Cfg.Port
	log.Printf("Starting server on port %s", port)
	router.Run(":" + port)
}