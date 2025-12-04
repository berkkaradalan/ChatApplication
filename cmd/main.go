package main

import (
	"log"

	"github.com/berkkaradalan/CoreGo"
	"github.com/berkkaradalan/CoreGo/auth"
	"github.com/berkkaradalan/CoreGo/database"
	"github.com/berkkaradalan/chatApp/config"
	"github.com/berkkaradalan/chatApp/handlers"
	"github.com/berkkaradalan/chatApp/middleware"
	"github.com/berkkaradalan/chatApp/repository"
	"github.com/berkkaradalan/chatApp/routes"
	"github.com/berkkaradalan/chatApp/services"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadEnv()

	core, err := corego.New(&corego.Config{
		Mongo: &database.MongoConfig{
			URL:      env.MONGODB_CONNECTION_URL,
			Database: env.MONGODB_DATABASE,
		},
		Auth: &auth.Config{
			Secret:       env.AUTH_SECRET,
			TokenExpiry:  env.TOKEN_EXPIRY,
			DatabaseName: env.AUTH_DATABASE,
		},
	})
	if err != nil {
		log.Fatal("Failed to initialize CoreGo:", err)
	}

	router := gin.Default()

	router.Use(middleware.CorsMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Server is running"})
	})

	router.POST("/signup", core.Auth.SignupHandler())
	router.POST("/login", core.Auth.LoginHandler())

	api := router.Group("/api")
	api.Use(core.Auth.Middleware())
	{
		api.GET("/profile", core.Auth.GetProfileHandler())
		api.PUT("/profile", core.Auth.UpdateProfileHandler())
		api.PUT("/password", core.Auth.ChangePasswordHandler())
		api.DELETE("/account", core.Auth.DeleteAccountHandler())
	}

	// Initialize repository, service, and handler
	roomRepo := repository.NewRoomRepository(core.Mongo.Collection("rooms"))
	roomService := services.NewRoomService(roomRepo)
	roomHandler := handlers.NewRoomHandler(roomService)
	messageRepo := repository.NewMessageRepository(core.Mongo.Collection("messages"))
	messageService := services.NewMessageService(messageRepo)
	messageHandler := handlers.NewMessageHandler(messageService)

	// Setup chat routes
	authConfig := config.NewAuthConfig(*env)
	routes.SetupRoutes(router, authConfig, roomHandler, messageHandler)

	log.Printf("Server starting on port %s...", env.PORT)
	if err := router.Run(":" + env.PORT); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
