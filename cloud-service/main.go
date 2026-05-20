package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	godotenv.Load()

	// 创建 Gin 路由
	router := gin.Default()

	// 注册路由
	registerAuthRoutes(router)
	registerSyncRoutes(router)
	registerAIRoutes(router)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func registerAuthRoutes(router *gin.Engine) {
	router.POST("/api/auth/device", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "device auth endpoint"})
	})
}

func registerSyncRoutes(router *gin.Engine) {
	router.POST("/api/sync/push", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "sync push endpoint"})
	})
	router.GET("/api/sync/pull", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "sync pull endpoint"})
	})
}

func registerAIRoutes(router *gin.Engine) {
	router.POST("/api/ai/chat", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ai chat endpoint"})
	})
}
