package api

import (
	"github.com/gin-gonic/gin"
)

// ChatRequest represents AI chat request
type ChatRequest struct {
	DeviceID  string `json:"device_id" binding:"required"`
	Message   string `json:"message" binding:"required"`
	Context   string `json:"context"`
}

// RegisterAIRoutes registers AI routes
func RegisterAIRoutes(router *gin.Engine) {
	router.POST("/api/ai/chat", handleAIChat)
}

func handleAIChat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement AI chat logic with SSE streaming
	c.JSON(200, gin.H{"message": "AI chat endpoint"})
}
