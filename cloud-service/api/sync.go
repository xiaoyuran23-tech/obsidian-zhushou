package api

import (
	"github.com/gin-gonic/gin"
)

// SyncPushRequest represents sync push request
type SyncPushRequest struct {
	DeviceID    string      `json:"device_id" binding:"required"`
	DeviceToken string      `json:"device_token" binding:"required"`
	Changes     []FileChange `json:"changes"`
}

// FileChange represents a file change
type FileChange struct {
	FilePath    string `json:"file_path" binding:"required"`
	Action      string `json:"action" binding:"required,oneof=create update delete"`
	Content     string `json:"content"`
	ContentHash string `json:"content_hash" binding:"required"`
	Timestamp   string `json:"timestamp" binding:"required"`
}

// SyncPullResponse represents sync pull response
type SyncPullResponse struct {
	Files       []SyncFile `json:"files"`
	SyncVersion string     `json:"sync_version"`
	NextSyncAfter string   `json:"next_sync_after"`
}

// SyncFile represents a file in sync response
type SyncFile struct {
	FilePath    string `json:"file_path"`
	Content     string `json:"content"`
	ContentHash string `json:"content_hash"`
	Timestamp   string `json:"timestamp"`
	HasConflict bool   `json:"has_conflict"`
}

// RegisterSyncRoutes registers sync routes
func RegisterSyncRoutes(router *gin.Engine) {
	router.POST("/api/sync/push", handleSyncPush)
	router.GET("/api/sync/pull", handleSyncPull)
}

func handleSyncPush(c *gin.Context) {
	var req SyncPushRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Process sync push
	c.JSON(200, gin.H{"sync_version": "v_1"})
}

func handleSyncPull(c *gin.Context) {
	since := c.Query("since")
	_ = since // TODO: Use since parameter for incremental sync

	c.JSON(200, SyncPullResponse{
		Files: []SyncFile{},
		SyncVersion: "v_1",
		NextSyncAfter: "TODO",
	})
}
