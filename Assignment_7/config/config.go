package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config struck for setting
type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPass     string
}

// Setting for global setting
func Setting() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf(".env file not found,this apps will use env in system")
	}

	config := Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_GORM_HOST"),
		DBPort:     os.Getenv("DB_GORM_PORT"),
		DBName:     os.Getenv("DB_GORM_NAME"),
		DBUsername: os.Getenv("DB_GORM_USERNAME"),
		DBPass:     os.Getenv("DB_GORM_PASS"),
	}

	return &config
}
