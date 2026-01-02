package setup

import (
	"fmt"

	"gorm.io/gorm"
)

func CleanDatabase(db *gorm.DB) error {
	if err := db.Exec("SET session_replication_role = 'replica';").Error; err != nil {
		return fmt.Errorf("failed to disable foreign keys: %w", err)
	}

	tables := []string{
		"team_knights",
		"knight_skill_levels",
		"knight_skills",
		"teams",
		"knights",
		"cosmo_attributes",
		"cosmos",
		"users",
	}

	for _, table := range tables {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
	}

	if err := db.Exec("SET session_replication_role = 'origin';").Error; err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	return nil
}
