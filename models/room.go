package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	RoomName	string				`json:"room_name" bson:"room_name"`
	LastMessage	string				`json:"last_message" bson:"last_name"`
	CreatedBy	string				`json:"created_by" bson:"created_by"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}

type CreateRoomRequest struct {
	RoomName	string				`json:"room_name" bson:"room_name"`
}

type ListRoomsRequest struct {
	Limit 		int		`json:"limit" bson:"limit"`
	Offset		int		`json:"offset" bson:"offset"`
	NewestFirst	bool	`json:"newest_first" bson:"newest_first"`
}

type ListRoomsResponse struct {
	RoomID				primitive.ObjectID	`json:"room_id" bson:"_id"`
	RoomName			string				`json:"room_name" bson:"room_name"`
	RoomLastMessage 	string				`json:"room_last_message" bson:"last_message"`
}