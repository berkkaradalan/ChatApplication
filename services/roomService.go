package services

import (
	"context"
	"errors"

	"github.com/berkkaradalan/chatApp/models"
	"github.com/berkkaradalan/chatApp/repository"
)


type RoomService struct {
	roomRepository 		*repository.RoomRepository
}

func NewRoomService(roomRepository *repository.RoomRepository) *RoomService {
	return &RoomService{
		roomRepository: roomRepository,	
	}
}

func (s *RoomService) GetRoom(ctx context.Context, roomID string) (*models.Room, error){
	if roomID == "" {
		return nil, errors.New("roomID can not be empty")
	}

	room, err := s.roomRepository.GetRoom(ctx, roomID)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomService) CreateRoom(ctx context.Context, roomName string, createdBy string) (*models.Room, error) {
	if roomName == "" {
		return nil, errors.New("roomName can not be empty")
	}

	room, err := s.roomRepository.CreateRoom(ctx, roomName, createdBy)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomService) ListRooms(ctx context.Context, limit int, offset int, newestFirst bool) (*[]models.ListRoomsResponse, error) {
	rooms, err := s.roomRepository.ListRooms(ctx, limit, offset, newestFirst)

	if err != nil {
		return nil, err
	}

	return rooms, nil
}