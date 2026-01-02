package setup

import (
	"saint-seiya-back/internal/infrastructure/database/entities"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.UserEntity{},
		&entities.KnightEntity{},
		&entities.KnightSkillEntity{},
		&entities.KnightSkillLevelEntity{},
		&entities.CosmoEntity{},
		&entities.CosmoAttributeEntity{},
		&entities.TeamEntity{},
		&entities.TeamKnightEntity{},
	)
}
