package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	RoomName	string				`json:"room_name" bson:"room_name"`
	LastMessage	string				`json:"last_name" bson:"last_name"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}