package handlers

import (
	"net/http"

	"github.com/berkkaradalan/chatApp/middleware"
	"github.com/berkkaradalan/chatApp/models"
	"github.com/berkkaradalan/chatApp/services"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	MessageService 		*services.MessageService
}

func NewMessageHandler (messageService *services.MessageService) *MessageHandler {
	return &MessageHandler{
		MessageService: messageService,
	}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req *models.SendMessageRequest

	if err := c.ShouldBind(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	claims := middleware.GetCurrentClaims(c)

	message, err := h.MessageService.SendMessage(c, req, claims)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (h *MessageHandler) GetMessages(c *gin.Context){
	var req *models.ListRoomMesaggesRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	chatMessages, err := h.MessageService.GetMessages(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": chatMessages,
		"count": len(chatMessages),
	})
}