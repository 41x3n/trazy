package controllers

import (
	"41x3n/trazy/config"
	"41x3n/trazy/models"
	"41x3n/trazy/utils"
)

var (
	infoLogger  = utils.InfoLogger
	errorLogger = utils.ErrorLogger
	cfg         models.Config
)

func init() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		errorLogger.Fatalf("Error loading configuration: %v", err)
	}
}
