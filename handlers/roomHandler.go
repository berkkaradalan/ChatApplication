package handlers

import (
	"net/http"

	"github.com/berkkaradalan/chatApp/middleware"
	"github.com/berkkaradalan/chatApp/models"
	"github.com/berkkaradalan/chatApp/services"
	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
    RoomService         *services.RoomService
}

func NewRoomHandler(roomService *services.RoomService) *RoomHandler {
	return &RoomHandler{
        RoomService: roomService,
	}
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
    roomID := c.Param("id")

    room, err := h.RoomService.GetRoom(c, roomID)

    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

    c.JSON(http.StatusOK, gin.H{
        "message": "room found",
        "room":    room,
    })
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var req *models.CreateRoomRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	claims := middleware.GetCurrentClaims(c)
	createdBy := claims.UserID

	room, err := h.RoomService.CreateRoom(c, req.RoomName, createdBy)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "room created successfuly!",
		"room": room,
	})
}

func (h *RoomHandler) ListRooms(c *gin.Context) {
	var req models.ListRoomsRequest

	_ = c.ShouldBindJSON(&req)

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	offset := req.Offset
	if offset < 0 {
		offset = 0
	}

	newestFirst := req.NewestFirst

	rooms, err := h.RoomService.ListRooms(c, limit, offset, newestFirst)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "rooms are listed",
		"rooms": rooms,
	})
}