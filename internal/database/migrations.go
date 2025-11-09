package database

import (
	"log"
	"saint-seiya-awakening/internal/models"
)

func MigrateDB() {
	log.Println("Running database migrations...")

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Knight{})
	DB.AutoMigrate(&models.Cosmo{})
	DB.AutoMigrate(&models.CosmoAttribute{})
	DB.AutoMigrate(&models.CosmoSubAttribute{})

	log.Println("Migrations completed successfully")
}
