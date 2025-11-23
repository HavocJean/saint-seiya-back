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
	DB.AutoMigrate(&models.KnightSkill{})
	DB.AutoMigrate(&models.KnightSkillLevel{})
	DB.AutoMigrate(&models.Team{})
	DB.AutoMigrate(&models.TeamKnight{})
	DB.AutoMigrate(&models.TeamVote{})
	DB.AutoMigrate(&models.TeamFavorite{})

	log.Println("Migrations completed successfully")
}
