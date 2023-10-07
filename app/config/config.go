package config

import (
	"41x3n/trazy/models"
	"41x3n/trazy/utils"
	"os"

	"github.com/dghubble/gologin/v2"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var errorLogger = utils.ErrorLogger

func LoadConfig() (models.Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		errorLogger.Fatalf("Error loading .env file: %v", err)
	}

	// Parse configuration values
	port := os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")
	githubClientID := os.Getenv("GITHUBCLIENTID")
	githubClientSecret := os.Getenv("GITHUBCLIENTSECRET")
	selfHost := os.Getenv("SELFHOST")
	jwtSecret := os.Getenv("JWTSECRET")

	githubOAuthConfig := &oauth2.Config{
		ClientID:     githubClientID,
		ClientSecret: githubClientSecret,
		RedirectURL:  selfHost + "/api/callback",
		Endpoint:     github.Endpoint,
	}

	stateConfig := gologin.DebugOnlyCookieConfig

	return models.Config{
		Port:                   port,
		GIN_MODE:               ginMode,
		GithubOAuthConfig:      githubOAuthConfig,
		GithubOAuthStateConfig: stateConfig,
		JwtSecret:              jwtSecret,
		SelfHost:               selfHost,
	}, nil
}
