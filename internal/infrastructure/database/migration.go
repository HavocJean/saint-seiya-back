package database

import (
	"log"
	"saint-seiya-back/internal/infrastructure/database/entities"
)

func MigrateDB() {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(&entities.UserEntity{})
	if err != nil {
		log.Fatalf("Migration failed: ", err)
	}

	log.Println("Migrations completed successfully")
}
