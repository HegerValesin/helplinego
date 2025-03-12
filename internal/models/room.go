package models

import "time"

type Room struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Number      string    `json:"number"`
	Description string    `json:"description"`
	FloorID     uint      `json:"floor_id"`
	Floor       Floor     `gorm:"foreignKey:FloorID" json:"floor,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
