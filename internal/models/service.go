package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ServiceStatus é um tipo para representar o status do serviço
type ServiceStatus string

const (
	Active   ServiceStatus = "ACTIVE"
	Inactive ServiceStatus = "INACTIVE"
	// Adicione outros status conforme necessário
)

type Services struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      ServiceStatus `json:"status" gorm:"type:varchar(20)"` // Ajuste o tamanho conforme necessário

	SectorID uuid.UUID `json:"sector_id"`
	Sector   Sector    `gorm:"foreignKey:SectorID" json:"sector"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeCreate é um hook do GORM que é chamado antes de um registro ser criado
func (s *Services) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	return
}

// BeforeUpdate é um hook do GORM que é chamado antes de um registro ser atualizado
func (s *Services) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}
