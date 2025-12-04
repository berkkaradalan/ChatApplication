package routes

import (
	"github.com/berkkaradalan/chatApp/config"
	"github.com/berkkaradalan/chatApp/handlers"
	"github.com/berkkaradalan/chatApp/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, authConfig *config.AuthConfig, roomHandler *handlers.RoomHandler, messageHandler *handlers.MessageHandler) {
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(authConfig))
	{
		protected.POST("/chat-room", roomHandler.CreateRoom)
		protected.GET("/chat-room/:id", roomHandler.GetRoom)
		protected.GET("/chat-rooms", roomHandler.ListRooms)
		protected.POST("/message", messageHandler.SendMessage)
		protected.GET("/message", messageHandler.GetMessages)
	}
}