package models

import "time"

type Facility struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Floors      []Floor   `gorm:"foreignKey:FacilityID" json:"floors,omitempty"`
}
