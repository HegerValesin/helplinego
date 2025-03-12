package models

import "time"

type Floor struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Number      int       `json:"number"`
	Description string    `json:"description"`
	FacilityID  uint      `json:"facility_id"`
	Facility    Facility  `gorm:"foreignKey:FacilityID" json:"facility,omitempty"`
	Rooms       []Room    `gorm:"foreignKey:FloorID" json:"rooms,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
