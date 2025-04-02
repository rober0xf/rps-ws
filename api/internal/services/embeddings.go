package services

import (
	"context"
	"rps/api"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room api.RoomDTO) (api.Room, error)
}

type roomService struct{}

func NewRoomSerice() RoomService {
	return roomService{}
}

func (r roomService) CreateRoom(ctx context.Context, room api.RoomDTO) (api.Room, error) {
	return api.Room{
		Name:        room.Name,
		MaxPlayers:  room.MaxPlayers,
		MinPlayers:  room.MinPlayers,
		Description: room.Description,
		Private:     room.Private,
	}, nil
}
