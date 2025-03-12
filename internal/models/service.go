package models

import "time"

type Service struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SectorID    uint      `json:"sector_id"`
	Sector      Sector    `gorm:"foreignKey:SectorID" json:"sector,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
