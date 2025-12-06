package handlers

import (
	"net/http"

	"github.com/berkkaradalan/chatApp/config"
	"github.com/berkkaradalan/chatApp/websocket"
	"github.com/gin-gonic/gin"
)

type WebSocketHandler struct {
	AuthConfig *config.AuthConfig
}

func NewWebSocketHandler(authConfig *config.AuthConfig) *WebSocketHandler {
	return &WebSocketHandler{
		AuthConfig: authConfig,
	}
}

// HandleRoomMessages upgrades HTTP connection to WebSocket for a specific room's messages
func (h *WebSocketHandler) HandleRoomMessages(c *gin.Context) {
	roomID := c.Query("roomId")
	token := c.Query("token")

	if roomID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomId is required"})
		return
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
		return
	}

	// Validate the token
	_, err := h.AuthConfig.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Validate that roomId is a valid MongoDB ObjectID format
	// This prevents using system channel names like "__system:rooms__"
	if len(roomID) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roomId format"})
		return
	}

	hub := websocket.GetHub()
	websocket.ServeWs(hub, c.Writer, c.Request, roomID)
}

// HandleRoomCreation upgrades HTTP connection to WebSocket for room creation events
func (h *WebSocketHandler) HandleRoomCreation(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
		return
	}

	// Validate the token
	_, err := h.AuthConfig.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	hub := websocket.GetHub()
	// Use a system channel name that cannot conflict with real room IDs
	websocket.ServeWs(hub, c.Writer, c.Request, "__system:rooms__")
}