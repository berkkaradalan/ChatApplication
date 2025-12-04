package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RoomID      primitive.ObjectID `json:"room_id" bson:"room_id"`
	SenderID    string             `json:"sender_id" bson:"sender_id"`
	SenderName  string             `json:"sender_name" bson:"sender_name"`
	MessageBody string             `json:"message_body" bson:"message_body"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type SendMessageRequest struct {
	ChatID 		string	`json:"chatID" bson:"chatID,omitempty"`
	MessageBody string	`json:"message_body" bson:"message_body"`
}