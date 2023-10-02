package config

import (
	"41x3n/trazy/utils"
	"os"

	"github.com/joho/godotenv"
)

var errorLogger = utils.ErrorLogger

type Config struct {
	// Add any configuration values you need here
	// For example, the port the server should run on
	Port     string
	GIN_MODE string
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		errorLogger.Fatalf("Error loading .env file: %v", err)
	}

	// Parse configuration values
	port := os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")

	return &Config{
		Port:     port,
		GIN_MODE: ginMode,
	}, nil
}
