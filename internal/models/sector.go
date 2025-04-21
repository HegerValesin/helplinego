package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sector struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeCreate is a GORM hook that is called before a record is created
func (s *Sector) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that is called before a record is updated
func (s *Sector) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
