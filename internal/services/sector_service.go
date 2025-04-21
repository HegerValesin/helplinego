package services

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/repository"
)

type SectorService struct {
	SectorRepo *repository.SectorRepository
}

func NewSectorService(sectorRepo *repository.SectorRepository) *SectorService {
	return &SectorService{SectorRepo: sectorRepo}
}

func (s *SectorService) Create(sector models.Sector) (*models.Sector, error) {
	err := s.SectorRepo.Create(&sector)
	if err != nil {
		return nil, err
	}
	return &sector, nil
}

func (s *SectorService) GetByID(id uuid.UUID) (*models.Sector, error) {
	return s.SectorRepo.FindByID(id)
}

func (s *SectorService) Update(sector models.Sector) (*models.Sector, error) {
	err := s.SectorRepo.Update(&sector)
	if err != nil {
		return nil, err
	}
	return &sector, nil
}

func (s *SectorService) Delete(id uuid.UUID) error {
	return s.SectorRepo.Delete(id)
}

func (s *SectorService) List(page, perPage int) ([]models.Sector, int64, error) {
	offset := (page - 1) * perPage
	return s.SectorRepo.List(perPage, offset)
}
