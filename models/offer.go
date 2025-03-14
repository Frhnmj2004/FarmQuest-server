package models

import (
	"time"
)

// Offer represents a special offer or promotion in the system
type Offer struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"size:1000" json:"description"`
	Type        string    `gorm:"size:50;not null" json:"type"` // discount, bonus, special
	Value       float64   `gorm:"not null" json:"value"`        // discount percentage or bonus points
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
