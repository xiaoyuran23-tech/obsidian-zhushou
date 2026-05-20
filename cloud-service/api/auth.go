package api

import (
	"github.com/gin-gonic/gin"
)

// PairDeviceRequest represents device pairing request
type PairDeviceRequest struct {
	Email      string `json:"email" binding:"required,email"`
	DeviceName string `json:"device_name" binding:"required"`
}

// PairDeviceResponse represents device pairing response
type PairDeviceResponse struct {
	DeviceCode string `json:"device_code"`
	ExpiresAt  string `json:"expires_at"`
}

// RegisterAuthRoutes registers authentication routes
func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/api/auth/device", handlePairDevice)
	router.POST("/api/auth/device/verify", handleVerifyDevice)
}

func handlePairDevice(c *gin.Context) {
	var req PairDeviceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// TODO: Generate device code and store in database
	c.JSON(200, PairDeviceResponse{
		DeviceCode: "TODO",
		ExpiresAt:  "TODO",
	})
}

func handleVerifyDevice(c *gin.Context) {
	// TODO: Verify device code
	c.JSON(200, gin.H{"message": "verify device endpoint"})
}
