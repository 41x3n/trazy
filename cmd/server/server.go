package main

import (
	"41x3n/tracy/config"
	"41x3n/tracy/routes"
	"41x3n/tracy/utils"
	"os"

	"github.com/gin-gonic/gin"
)

var infoLogger = utils.InfoLogger
var errorLogger = utils.ErrorLogger

func get_port(port string) string {
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		errorLogger.Fatalf("Error loading configuration: %v", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.GIN_MODE)

	// Create custom router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	port := get_port(cfg.Port)

	// Start server
	infoLogger.Fatal(router.Run(port))
}
