package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name       string    `json:"name"`
	RoomNumber string    `json:"room_number"`
	Area       float64   `json:"area"`

	FloorID uuid.UUID `json:"floor_id"`
	Floor   Floor     `gorm:"foreignKey:FloorID" json:"floor"`

	SectorID uuid.UUID `json:"sector_id"`
	Sector   Sector    `gorm:"foreignKey:SectorID" json:"sector"`

	CreatedByID uuid.UUID `json:"created_by"`
	CreatedBy   User      `gorm:"foreignKey:CreatedByID" json:"created_by"`

	UpdatedByID uuid.UUID `json:"updated_by"`
	UpdatedBy   User      `gorm:"foreignKey:UpdatedByID" json:"updated_by"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeCreate is a GORM hook that is called before a record is created
func (r *Room) BeforeCreate(tx *gorm.DB) (err error) {
	r.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that is called before a record is updated
func (r *Room) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdatedAt = time.Now()
	return
}
