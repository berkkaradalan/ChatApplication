package repository

import (
	"context"
	"time"

	"github.com/berkkaradalan/chatApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoomRepository struct {
	mongoDB				*mongo.Collection
}

func NewRoomRepository (mongo *mongo.Collection) *RoomRepository{
	return &RoomRepository{
		mongoDB: mongo,
	}
}

func (r *RoomRepository) GetRoom(ctx context.Context, roomID string) (*models.Room, error){
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, err
	}

	var room models.Room
	err = r.mongoDB.FindOne(ctx, bson.M{"_id": objectID}).Decode(&room)

	return &room, err
}

func (r *RoomRepository) CreateRoom(ctx context.Context, roomName string, createdBy string) (*models.Room, error) {
	room := &models.Room{
		ID:   primitive.NewObjectID(),
		RoomName: roomName,
		LastMessage: "",
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := r.mongoDB.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *RoomRepository) ListRooms(ctx context.Context, limit int, offset int, newestFirst bool) (*[]models.ListRoomsResponse, error) {
	var rooms []models.ListRoomsResponse

	filter := bson.M{}

	sortOrder := -1
	if !newestFirst {
		sortOrder = 1
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: sortOrder}})
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	cursor, err := r.mongoDB.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return &rooms, nil
}