package repository

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"gorm.io/gorm"
)

type FacilityRepository struct {
	DB *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) *FacilityRepository {
	return &FacilityRepository{DB: db}
}

func (r *FacilityRepository) Create(facility *models.Facility) error {
	return r.DB.Create(facility).Error
}

func (r *FacilityRepository) FindByID(id uuid.UUID) (*models.Facility, error) {
	var facility models.Facility
	err := r.DB.Preload("Floors").First(&facility, id).Error
	if err != nil {
		return nil, err
	}
	return &facility, nil
}

func (r *FacilityRepository) Update(facility *models.Facility) error {
	return r.DB.Save(facility).Error
}

func (r *FacilityRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Facility{}, id).Error
}

func (r *FacilityRepository) List(limit, offset int) ([]models.Facility, int64, error) {
	var facilities []models.Facility
	var count int64

	if err := r.DB.Model(&models.Facility{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Limit(limit).Offset(offset).Find(&facilities).Error; err != nil {
		return nil, 0, err
	}

	return facilities, count, nil
}
