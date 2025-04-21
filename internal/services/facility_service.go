package services

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/repository"
)

type FacilityService struct {
	FacilityRepo *repository.FacilityRepository
}

func NewFacilityService(facilityRepo *repository.FacilityRepository) *FacilityService {
	return &FacilityService{FacilityRepo: facilityRepo}
}

func (s *FacilityService) Create(facility models.Facility) (*models.Facility, error) {
	err := s.FacilityRepo.Create(&facility)
	if err != nil {
		return nil, err
	}
	return &facility, nil
}

func (s *FacilityService) GetByID(id uuid.UUID) (*models.Facility, error) {
	return s.FacilityRepo.FindByID(id)
}

func (s *FacilityService) Update(facility models.Facility) (*models.Facility, error) {
	err := s.FacilityRepo.Update(&facility)
	if err != nil {
		return nil, err
	}
	return &facility, nil
}

func (s *FacilityService) Delete(id uuid.UUID) error {
	return s.FacilityRepo.Delete(id)
}

func (s *FacilityService) List(page, perPage int) ([]models.Facility, int64, error) {
	offset := (page - 1) * perPage
	return s.FacilityRepo.List(perPage, offset)
}
