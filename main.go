package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"web/configs"
	"web/routes"
)

func main() {
	// Create context for database operations
	ctx := context.Background()

	// Initialize database connection
	configs.InitDB()
	defer configs.CloseDB(ctx)

	// Set up Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Define home route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id": 1,
			"message": "Welcome to the Backend Server of GoLang🐰",
		})
	})

	// Setup API routes
	routes.SetupRoutes(r)

	// Start the server
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	log.Printf("Server running on port %s 🚀", serverPort)
	if err := r.Run(":" + serverPort); err != nil {
		log.Fatal("Server failed:", err)
	}
}
