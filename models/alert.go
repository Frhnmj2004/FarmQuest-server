package models

import "time"

// Alert represents a notification or alert for a user
type Alert struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Message   string    `gorm:"size:1000;not null" json:"message"`
	Type      string    `gorm:"size:50;not null" json:"type"` // info, warning, success
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
