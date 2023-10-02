package controllers

import (
	"41x3n/trazy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var infoLogger = utils.InfoLogger

const gifImage = "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\x00\x00\xff\xff\xff\x00\x00\x00\x2c\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3b"

func Pixel(c *gin.Context) {
	id := c.Query("id")
	c.Header("Content-Type", "image/gif")

	if id == "" {
		infoLogger.Println("Pixel request received from unknown user", c.ClientIP())
		c.Data(http.StatusOK, "image/gif", []byte(gifImage))
		return
	}

	reqInfo, err := utils.GetInfoFromRequest(c)
	if err != nil {
		infoLogger.Println("Pixel request received from user", id, "but no info was found")
		c.Data(http.StatusOK, "image/gif", []byte(gifImage))
		return
	}

	infoLogger.Printf("Pixel request received from user %s  %+v\n", id, reqInfo)
	c.Data(http.StatusOK, "image/gif", []byte(gifImage))
}
