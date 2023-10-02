package utils

import (
	"41x3n/trazy/models"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func GetInfoFromRequest(c *gin.Context) (models.RequestInfo, error) {
	if c.ClientIP() == "" {
		return models.RequestInfo{}, errors.New("no IP address found")
	}

	reqInfo := models.RequestInfo{
		IPAddress:      c.ClientIP(),
		UserAgent:      c.GetHeader("User-Agent"),
		Referrer:       c.GetHeader("Referer"),
		Timestamp:      time.Now().UTC().Format(time.RFC3339),
		AcceptLanguage: c.GetHeader("Accept-Language"),
		DoNotTrack:     c.GetHeader("DNT"),
	}

	return reqInfo, nil
}
