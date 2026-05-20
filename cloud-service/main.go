package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/obsidian-zhushou/cloud-service/api"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Create Gin router
	router := gin.Default()

	// Register routes
	api.RegisterAuthRoutes(router)
	api.RegisterSyncRoutes(router)
	api.RegisterAIRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting OBzhushou cloud service on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
