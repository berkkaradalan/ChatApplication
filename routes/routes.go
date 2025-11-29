package routes

import (
	"github.com/berkkaradalan/chatApp/config"
	"github.com/berkkaradalan/chatApp/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authConfig *config.AuthConfig) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CorsMiddleware())

	router.GET("/message/:id")

	protected := router.Use(middleware.AuthMiddleware(authConfig))
	{
		protected.POST("/message")
	}

	return router
}