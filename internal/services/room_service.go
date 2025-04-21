// internal/services/room_service.go
package services

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/repository"
)

type RoomService struct {
	RoomRepo  *repository.RoomRepository
	FloorRepo *repository.FloorRepository
}

func NewRoomService(roomRepo *repository.RoomRepository, floorRepo *repository.FloorRepository) *RoomService {
	return &RoomService{
		RoomRepo:  roomRepo,
		FloorRepo: floorRepo,
	}
}

func (s *RoomService) Create(room models.Room) (*models.Room, error) {
	// Verificar se o andar existe
	_, err := s.FloorRepo.FindByID(room.FloorID)
	if err != nil {
		return nil, err
	}

	err = s.RoomRepo.Create(&room)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (s *RoomService) GetByID(id uuid.UUID) (*models.Room, error) {
	return s.RoomRepo.FindByID(id)
}

func (s *RoomService) Update(room models.Room) (*models.Room, error) {
	// Verificar se o andar existe se o ID do andar foi alterado
	if room.FloorID != [16]byte{} {
		_, err := s.FloorRepo.FindByID(room.FloorID)
		if err != nil {
			return nil, err
		}
	}

	err := s.RoomRepo.Update(&room)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (s *RoomService) Delete(id uuid.UUID) error {
	return s.RoomRepo.Delete(id)
}

func (s *RoomService) List(page, perPage int) ([]models.Room, int64, error) {
	offset := (page - 1) * perPage
	return s.RoomRepo.List(perPage, offset)
}

func (s *RoomService) ListByFloor(floorID uuid.UUID) ([]models.Room, error) {
	return s.RoomRepo.FindByFloorID(floorID)
}
