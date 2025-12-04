package services

import (
	"context"

	"github.com/berkkaradalan/chatApp/config"
	"github.com/berkkaradalan/chatApp/models"
	"github.com/berkkaradalan/chatApp/repository"
)

type MessageService struct {
	messageRepository 		*repository.MessageRepository
}

func NewMessageService (messageRepository *repository.MessageRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
	}
}

func (s *MessageService) SendMessage(ctx context.Context, requestBody *models.SendMessageRequest, claims *config.JWTClaims) (*models.Message, error){
	message, err := s.messageRepository.SendMessage(ctx, requestBody, claims)

	if err != nil {
		return nil, err
	}

	return message, nil
}
