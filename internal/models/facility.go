package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Facility struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string    `json:"name" form:"name"`
	Abreviation string    `json:"abreviation" form:"abreviation"`

	// Relacionamentos
	Floors []Floor `gorm:"foreignKey:FacilityID" json:"floors,omitempty"`

	// Campos de auditoria
	CreatedBy   *User      `json:"created_by,omitempty"`
	CreatedByID uuid.UUID  `gorm:"type:uuid;column:created_by" json:"created_by_id"`
	UpdatedBy   *User      `json:"updated_by,omitempty"`
	UpdatedByID uuid.UUID  `gorm:"type:uuid;column:updated_by" json:"updated_by_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
}

func (f *Facility) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	f.CreatedAt = time.Now()
	return
}

func (Facility) TableName() string {
	return "facility"
}

func (f *Facility) BeforeUpdate(tx *gorm.DB) (err error) {
	f.UpdatedAt = time.Now()
	return
}
