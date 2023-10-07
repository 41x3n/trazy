package models

import (
	"github.com/dghubble/gologin/v2"
	"golang.org/x/oauth2"
)

type Config struct {
	// Add any configuration values you need here
	// For example, the port the server should run on
	Port                   string
	GIN_MODE               string
	GithubOAuthConfig      *oauth2.Config
	GithubOAuthStateConfig gologin.CookieConfig
	JwtSecret              string
	SelfHost               string
}
