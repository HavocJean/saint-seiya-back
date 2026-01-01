package database

import (
	"fmt"
	"log"
	"saint-seiya-back/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Cfg.DBHost,
		config.Cfg.DBUser,
		config.Cfg.DBPass,
		config.Cfg.DBName,
		config.Cfg.DBPort,
	)

	var db *gorm.DB
	var err error
	maxAttempts := 10

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			log.Println("Database connection established")
			return
		}

		wait := time.Duration(attempt*500) * time.Millisecond
		log.Printf("Database not ready (attempt %d/%d): %v — retrying in %s", attempt, maxAttempts, err, wait)
		time.Sleep(wait)
	}

	log.Fatalf("Failed to connect to database after %d attempts: %v", maxAttempts, err)
}

func GetDB() *gorm.DB {
	return DB
}
