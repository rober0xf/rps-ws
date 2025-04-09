package services

import (
	"context"
	"github.com/uptrace/bun"
	"rps/pkg/models"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room models.RoomDTO) (models.Room, error)
	GetRoom(ctx context.Context, id uint64) (models.Room, error)
	ListRooms(ctx context.Context) ([]models.Room, error)
}

type roomServiceImpl struct {
	db *bun.DB
}

func (s *roomServiceImpl) GenerateAddress() (string, error) {
	return "todo", nil
}

// SERVICES THAT WILL BE USED BY THE HANDLERS

func (s *roomServiceImpl) CreateRoomService(ctx context.Context, dto models.RoomDTO) (models.Room, error) {
	address, err := s.GenerateAddress()
	if err != nil {
		return models.Room{}, err
	}

	room := models.Room{
		Name:        dto.Name,
		MaxPlayers:  dto.MaxPlayers,
		Description: dto.Description,
		Private:     dto.Private,
		Address:     address,
	}

	_, err = s.db.NewInsert().Model(&room).Exec(ctx)
	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}

func (s *roomServiceImpl) GetRoomByAddressService(ctx context.Context, address string) (models.Room, error) {
	room := models.Room{}

	err := s.db.NewSelect().Model(&room).Where("address = ?", address).Scan(ctx)
	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}

func (s *roomServiceImpl) ListRoomsService(ctx context.Context) ([]models.Room, error) {
	rooms := []models.Room{}

	err := s.db.NewSelect().Model(&rooms).Where("private = ?", false).Scan(ctx)
	if err != nil {
		// we can return nil because in slices nil means that there is not a slice
		return nil, err
	}

	return rooms, nil
}
