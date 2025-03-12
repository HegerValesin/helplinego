package models

import "time"

type TicketStatus string

const (
	StatusOpen       TicketStatus = "open"
	StatusInProgress TicketStatus = "in_progress"
	StatusResolved   TicketStatus = "resolved"
	StatusClosed     TicketStatus = "closed"
)

type Ticket struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      TicketStatus `json:"status"`
	Priority    int          `json:"priority"` // 1-5, com 5 sendo a mais alta
	UserID      uint         `json:"user_id"`
	User        User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ServiceID   uint         `json:"service_id"`
	Service     Service      `gorm:"foreignKey:ServiceID" json:"service,omitempty"`
	RoomID      uint         `json:"room_id"`
	Room        Room         `gorm:"foreignKey:RoomID" json:"room,omitempty"`
	AssignedTo  uint         `json:"assigned_to"`
	Assignee    User         `gorm:"foreignKey:AssignedTo" json:"assignee,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	ClosedAt    *time.Time   `json:"closed_at"`
}
