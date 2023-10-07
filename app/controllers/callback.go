package controllers

import (
	"41x3n/trazy/models"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/dghubble/gologin/v2/github"
	"github.com/dghubble/gologin/v2/oauth2"
	"github.com/gin-gonic/gin"
)

const (
	cookieName = "jwt"
)

func Callback(c *gin.Context) {
	stateConfig := cfg.GithubOAuthStateConfig
	config := cfg.GithubOAuthConfig

	defer func() {
		if err := recover(); err != nil {
			errorLogger.Printf("Error verifying the auth token: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()

	github.StateHandler(stateConfig, github.CallbackHandler(config, issueSession(), nil)).ServeHTTP(c.Writer, c.Request)

}

// issueSession returns an http.Handler that issues a session for the authenticated user.
func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// Get the OAuth token and GitHub user.
		token, err := oauth2.TokenFromContext(ctx)
		if err != nil {
			errorLogger.Printf("Error getting token from context: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			errorLogger.Printf("Error getting user from context: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Handle errors gracefully.
		if githubUser == nil {
			err := errors.New("GitHub user not found")
			errorLogger.Printf("Error issuing session: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Grant the visitor a session (cookie, token, etc.).
		infoLogger.Printf("User - %s is authenticated: %v\n", *githubUser.Email, token.Valid())

		// Create a UserInfo struct to hold the user information and access token.
		userInfo := models.UserInfo{
			Email:     *githubUser.Email,
			Name:      *githubUser.Name,
			AvatarURL: *githubUser.AvatarURL,
		}

		// Create a new JWT token with the user information.
		claims := jwt.MapClaims{
			"user_info":    userInfo,
			"access_token": token.AccessToken,
			"exp":          time.Now().Add(time.Hour * 240).Unix(),
		}
		jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		jwt_secret := cfg.JwtSecret
		signedToken, err := jwt_token.SignedString([]byte(jwt_secret))
		if err != nil {
			errorLogger.Printf("Error signing JWT: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the access token in the response header.
		w.Header().Set("X-Access-Token", token.AccessToken)

		// Set the JWT as a cookie in the response.
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    signedToken,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})

		redirectURL := getRedirectURL(cfg)

		// Redirect the user to the frontend web app.
		http.Redirect(w, req, redirectURL, http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

func getRedirectURL(cfg models.Config) string {
	return cfg.SelfHost + "/api/health"
}
