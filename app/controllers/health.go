package controllers

import (
	"41x3n/trazy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	message := models.Message{Message: "OK"}
	c.JSON(http.StatusOK, message)
}
