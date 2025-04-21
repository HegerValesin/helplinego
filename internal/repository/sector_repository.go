package repository

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"gorm.io/gorm"
)

type SectorRepository struct {
	DB *gorm.DB
}

func NewSectorRepository(db *gorm.DB) *SectorRepository {
	return &SectorRepository{DB: db}
}

func (r *SectorRepository) Create(sector *models.Sector) error {
	return r.DB.Create(sector).Error
}

func (r *SectorRepository) FindByID(id uuid.UUID) (*models.Sector, error) {
	var sector models.Sector
	err := r.DB.First(&sector, id).Error
	if err != nil {
		return nil, err
	}
	return &sector, nil
}

func (r *SectorRepository) Update(sector *models.Sector) error {
	return r.DB.Save(sector).Error
}

func (r *SectorRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Sector{}, id).Error
}

func (r *SectorRepository) List(limit, offset int) ([]models.Sector, int64, error) {
	var sectors []models.Sector
	var count int64

	if err := r.DB.Model(&models.Sector{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Limit(limit).Offset(offset).Find(&sectors).Error; err != nil {
		return nil, 0, err
	}

	return sectors, count, nil
}
