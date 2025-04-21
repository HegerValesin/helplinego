package repository

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"gorm.io/gorm"
)

type FloorRepository struct {
	DB *gorm.DB
}

func NewFloorRepository(db *gorm.DB) *FloorRepository {
	return &FloorRepository{DB: db}
}

func (r *FloorRepository) Create(floor *models.Floor) error {
	return r.DB.Create(floor).Error
}

func (r *FloorRepository) FindByID(id uuid.UUID) (*models.Floor, error) {
	var floor models.Floor
	err := r.DB.Preload("Facility").Preload("Rooms").First(&floor, id).Error
	if err != nil {
		return nil, err
	}
	return &floor, nil
}

func (r *FloorRepository) Update(floor *models.Floor) error {
	return r.DB.Save(floor).Error
}

func (r *FloorRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Floor{}, id).Error
}

func (r *FloorRepository) List(limit, offset int) ([]models.Floor, int64, error) {
	var floors []models.Floor
	var count int64

	if err := r.DB.Model(&models.Floor{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Preload("Facility").Limit(limit).Offset(offset).Find(&floors).Error; err != nil {
		return nil, 0, err
	}

	return floors, count, nil
}

func (r *FloorRepository) FindByFacilityID(facilityID uuid.UUID) ([]models.Floor, error) {
	var floors []models.Floor
	err := r.DB.Where("facility_id = ?", facilityID).Preload("Rooms").Find(&floors).Error
	if err != nil {
		return nil, err
	}
	return floors, nil
}
