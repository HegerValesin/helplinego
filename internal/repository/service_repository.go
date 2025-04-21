package repository

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	DB *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{DB: db}
}

func (r *ServiceRepository) Create(service *models.Services) error {
	return r.DB.Create(service).Error
}

func (r *ServiceRepository) FindByID(id uuid.UUID) (*models.Services, error) {
	var service models.Services
	err := r.DB.Preload("Sector").First(&service, id).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *ServiceRepository) Update(service *models.Services) error {
	return r.DB.Save(service).Error
}

func (r *ServiceRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Services{}, id).Error
}

func (r *ServiceRepository) List(limit, offset int) ([]models.Services, int64, error) {
	var services []models.Services
	var count int64

	if err := r.DB.Model(&models.Services{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Preload("Sector").Limit(limit).Offset(offset).Find(&services).Error; err != nil {
		return nil, 0, err
	}

	return services, count, nil
}

func (r *ServiceRepository) FindBySectorID(sectorID uuid.UUID, limit, offset int) ([]models.Services, int64, error) {
	var services []models.Services
	var count int64

	if err := r.DB.Model(&models.Services{}).Where("sector_id = ?", sectorID).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Where("sector_id = ?", sectorID).Preload("Sector").Limit(limit).Offset(offset).Find(&services).Error; err != nil {
		return nil, 0, err
	}

	return services, count, nil
}
