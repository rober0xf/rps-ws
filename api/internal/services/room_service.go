package services

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"math/rand"
	"rps/pkg/models"
)

type RoomService interface {
	CreateRoomService(ctx context.Context, room models.RoomDTO) (models.Room, error)
	GetRoomByIDService(ctx context.Context, id uint64) (models.Room, error)
	GetRoomByAddressService(ctx context.Context, address string) (models.Room, error)
	ListRoomsService(ctx context.Context) ([]models.Room, error)
}

type roomServiceImpl struct {
	db *bun.DB
}

// address for private rooms
func (s *roomServiceImpl) GenerateAddress(length int) (string, error) {
	if length < 0 || length > 128 {
		return "", fmt.Errorf("length must be between 0 and 128, current: %d", length)
	}

	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b), nil
}

// SERVICES THAT WILL BE USED BY THE HANDLERS

func (s *roomServiceImpl) CreateRoomService(ctx context.Context, dto models.RoomDTO) (models.Room, error) {
	address := ""
	if dto.Private {
		a, err := s.GenerateAddress(50)
		if err != nil {
			return models.Room{}, fmt.Errorf("failed to generate private address: %w", err)
		}
		address = a
	}

	room := models.Room{
		Name:        dto.Name,
		MaxPlayers:  dto.MaxPlayers,
		Description: dto.Description,
		Private:     dto.Private,
		Address:     address,
	}

	_, err := s.db.NewInsert().Model(&room).Exec(ctx)
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

func (s *roomServiceImpl) GetRoomByIDService(ctx context.Context, id uint64) (models.Room, error) {
	room := models.Room{}

	err := s.db.NewSelect().Model(&room).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}
