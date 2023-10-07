package controllers

import (
	"net/http"

	"github.com/dghubble/gologin/v2/github"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	stateConfig := cfg.GithubOAuthStateConfig
	config := cfg.GithubOAuthConfig

	defer func() {
		if err := recover(); err != nil {
			errorLogger.Printf("Error serving Github Auth page: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()

	github.StateHandler(stateConfig, github.LoginHandler(config, nil)).ServeHTTP(c.Writer, c.Request)
}
