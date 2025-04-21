package repository

import (
	"github.com/google/uuid"
	"github.com/hegervalesin/helplinego/internal/models"
	"gorm.io/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{DB: db}
}

func (r *RoomRepository) Create(room *models.Room) error {
	return r.DB.Create(room).Error
}

func (r *RoomRepository) FindByID(id uuid.UUID) (*models.Room, error) {
	var room models.Room
	err := r.DB.Preload("Floor").Preload("Floor.Facility").First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) Update(room *models.Room) error {
	return r.DB.Save(room).Error
}

func (r *RoomRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Room{}, id).Error
}

func (r *RoomRepository) List(limit, offset int) ([]models.Room, int64, error) {
	var rooms []models.Room
	var count int64

	if err := r.DB.Model(&models.Room{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Preload("Floor").Limit(limit).Offset(offset).Find(&rooms).Error; err != nil {
		return nil, 0, err
	}

	return rooms, count, nil
}

func (r *RoomRepository) FindByFloorID(floorID uuid.UUID) ([]models.Room, error) {
	var rooms []models.Room
	err := r.DB.Where("floor_id = ?", floorID).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
