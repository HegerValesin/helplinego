package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Floor struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `json:"name"`
	FloorNumber string    `json:"floor_number"`
	FacilityID  uuid.UUID `json:"facility_id"`
	Facility    Facility  `gorm:"foreignKey:FacilityID" json:"facility"`
	Rooms       []Room    `gorm:"foreignKey:FloorID" json:"rooms"`
	CreatedByID uuid.UUID `json:"created_by"`
	CreatedBy   User      `gorm:"foreignKey:CreatedByID" json:"created_by"`
	UpdatedByID uuid.UUID `json:"updated_by"`
	UpdatedBy   User      `gorm:"foreignKey:UpdatedByID" json:"updated_by"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeCreate is a GORM hook that is called before a record is created
func (f *Floor) BeforeCreate(tx *gorm.DB) (err error) {
	f.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that is called before a record is updated
func (f *Floor) BeforeUpdate(tx *gorm.DB) (err error) {
	f.UpdatedAt = time.Now()
	return
}
