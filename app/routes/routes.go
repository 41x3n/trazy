package routes

import (
	"41x3n/trazy/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/health", controllers.Health)
	api.GET("/pixel", controllers.Pixel)
	api.GET("/login", controllers.Login)
	api.GET("/callback", controllers.Callback)
}
