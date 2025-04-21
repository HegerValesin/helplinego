package services

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"github.com/hegervalesin/helplinego/internal/repository"
)

type ServiceService struct {
	ServiceRepo *repository.ServiceRepository
	SectorRepo  *repository.SectorRepository
}

func NewServiceService(serviceRepo *repository.ServiceRepository, sectorRepo *repository.SectorRepository) *ServiceService {
	return &ServiceService{
		ServiceRepo: serviceRepo,
		SectorRepo:  sectorRepo,
	}
}

func (s *ServiceService) Create(service models.Services) (*models.Services, error) {
	// Verificar se o setor existe
	_, err := s.SectorRepo.FindByID(service.SectorID)
	if err != nil {
		return nil, err
	}

	err = s.ServiceRepo.Create(&service)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *ServiceService) GetByID(id uuid.UUID) (*models.Services, error) {
	return s.ServiceRepo.FindByID(id)
}

func (s *ServiceService) Update(service models.Services) (*models.Services, error) {
	// Verificar se o setor existe se o ID do setor foi alterado
	if service.SectorID != uuid.Nil {
		_, err := s.SectorRepo.FindByID(service.SectorID)
		if err != nil {
			return nil, err
		}
	}

	err := s.ServiceRepo.Update(&service)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *ServiceService) Delete(id uuid.UUID) error {
	return s.ServiceRepo.Delete(id)
}

func (s *ServiceService) List(page, perPage int) ([]models.Services, int64, error) {
	offset := (page - 1) * perPage
	return s.ServiceRepo.List(perPage, offset)
}

func (s *ServiceService) ListBySector(sectorID uuid.UUID, page, perPage int) ([]models.Services, int64, error) {
	offset := (page - 1) * perPage
	return s.ServiceRepo.FindBySectorID(sectorID, perPage, offset)
}
