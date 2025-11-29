package database

import (
	"log"
	"saint-seiya-awakening/internal/models"
)

func MigrateDB() {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %w", err)
	}

	log.Println("Migrations completed successfully")
}
