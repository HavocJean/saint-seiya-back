package database

import (
	"log"
	"saint-seiya-back/internal/infrastructure/database/entities"
)

func MigrateDB() {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&entities.UserEntity{},
		&entities.KnightEntity{},
		&entities.CosmoEntity{},
		&entities.CosmoAttributeEntity{},
		&entities.TeamEntity{},
		&entities.TeamKnightEntity{},
	)
	if err != nil {
		log.Fatalf("Migration failed: ", err)
	}

	log.Println("Migrations completed successfully")
}
