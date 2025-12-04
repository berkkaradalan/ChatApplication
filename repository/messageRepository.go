package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/berkkaradalan/chatApp/config"
	"github.com/berkkaradalan/chatApp/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	mongoDB 		*mongo.Collection
}

func NewMessageRepository(mongo *mongo.Collection) *MessageRepository {
	return &MessageRepository{
		mongoDB: mongo,
	}
}

func(r *MessageRepository) SendMessage(ctx context.Context, requestBody *models.SendMessageRequest, claims *config.JWTClaims) (*models.Message, error) {
	roomID, err := primitive.ObjectIDFromHex(requestBody.ChatID)
	if err != nil {
		return nil, fmt.Errorf("invalid room ID format: %w", err)
	}

	message := models.Message{
		ID:				primitive.NewObjectID(),
		RoomID:     	roomID,
		SenderID:   	claims.UserID,
		SenderName: 	claims.UserName,
		MessageBody:	requestBody.MessageBody,
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}

	insertResult, err := r.mongoDB.InsertOne(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("failed to insert message: %s", err)
	}

	if insertResult.InsertedID == nil {
		return nil, fmt.Errorf("message insertion failed: no ID returned")
	}

	return &message, nil
}
