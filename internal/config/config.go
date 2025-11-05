package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port		string
	DBHost		string
	DBPort		string
	DBUser		string
	DBPass		string
	DBName		string
}

var Cfg Config

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Cfg = Config{
		Port:   os.Getenv("PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}
}

func Getenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}