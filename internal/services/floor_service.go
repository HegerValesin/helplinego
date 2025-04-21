package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/repository"
)

type FloorService struct {
	FloorRepo    *repository.FloorRepository
	FacilityRepo *repository.FacilityRepository
}

func NewFloorService(floorRepo *repository.FloorRepository, facilityRepo *repository.FacilityRepository) *FloorService {
	return &FloorService{
		FloorRepo:    floorRepo,
		FacilityRepo: facilityRepo,
	}
}

func (s *FloorService) Create(floor models.Floor) (*models.Floor, error) {
	// Verificar se a facility existe
	facilityUUID := uuid.New()
	err := facilityUUID.UnmarshalText([]byte(fmt.Sprintf("%d", floor.FacilityID)))
	if err != nil {
		return nil, err
	}
	_, err = s.FacilityRepo.FindByID(facilityUUID)
	if err != nil {
		return nil, err
	}

	err = s.FloorRepo.Create(&floor)
	if err != nil {
		return nil, err
	}
	return &floor, nil
}

func (s *FloorService) GetByID(id uuid.UUID) (*models.Floor, error) {
	return s.FloorRepo.FindByID(id)
}

func (s *FloorService) Update(floor models.Floor) (*models.Floor, error) {
	// Verificar se a facility existe se o ID da facility foi alterado
	if floor.FacilityID != [16]byte{} {
		facilityUUID := uuid.New()
		err := facilityUUID.UnmarshalText([]byte(fmt.Sprintf("%d", floor.FacilityID)))
		if err != nil {
			return nil, err
		}
		_, err = s.FacilityRepo.FindByID(facilityUUID)
		if err != nil {
			return nil, err
		}
	}

	err := s.FloorRepo.Update(&floor)
	if err != nil {
		return nil, err
	}
	return &floor, nil
}

func (s *FloorService) Delete(id uuid.UUID) error {
	return s.FloorRepo.Delete(id)
}

func (s *FloorService) List(page, perPage int) ([]models.Floor, int64, error) {
	offset := (page - 1) * perPage
	return s.FloorRepo.List(perPage, offset)
}

func (s *FloorService) ListByFacility(facilityID uuid.UUID) ([]models.Floor, error) {
	return s.FloorRepo.FindByFacilityID(facilityID)
}
