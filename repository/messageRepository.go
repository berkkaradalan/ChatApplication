package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/berkkaradalan/chatApp/config"
	"github.com/berkkaradalan/chatApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func(r *MessageRepository) GetMessages(ctx context.Context, roomID string, limit int, offset int, newestFirst bool) ([]*models.Message, error){
	var messages []*models.Message

	objRoomID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, fmt.Errorf("invalid room ID format: %w", err)
	}

	filter := bson.M{"room_id": objRoomID}

	sortOrder := -1
	if !newestFirst {
		sortOrder = 1
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: sortOrder}}).
		SetSkip(int64(offset)).
		SetLimit(int64(limit))

	cursor, err := r.mongoDB.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &messages); err != nil {
		return nil, fmt.Errorf("failed to decode messages: %w", err)
	}

	return messages, nil
}