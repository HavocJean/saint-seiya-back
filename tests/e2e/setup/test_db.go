package setup

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	TestDBHost     = "localhost"
	TestDBPort     = "5433"
	TestDBUser     = "postgres"
	TestDBPassword = "postgres"
	TestDBName     = "saintseiya_test"
)

func ConnectTestDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		TestDBHost,
		TestDBUser,
		TestDBPassword,
		TestDBName,
		TestDBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database: %w", err)
	}

	return db, nil
}
