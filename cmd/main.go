package main

import (
	"log"
	"os"

	"github.com/berkkaradalan/CoreGo"
	"github.com/berkkaradalan/CoreGo/auth"
	"github.com/berkkaradalan/CoreGo/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	core, err := corego.New(&corego.Config{
		Mongo: &database.MongoConfig{
			URL:      getEnv("MONGODB_CONNECTION_URL", "mongodb://localhost:27017"),
			Database: getEnv("MONGODB_DATABASE", "chatApplication"),
		},
		Auth: &auth.Config{
			Secret:       getEnv("AUTH_SECRET", "default-secret-key"),
			TokenExpiry:  1440,
			DatabaseName: getEnv("AUTH_DATABASE", "users"),
		},
	})
	if err != nil {
		log.Fatal("Failed to initialize CoreGo:", err)
	}

	router := gin.Default()

	router.Use(corsMiddleware())

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

	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
