package models

import (
	"time"
)

// Order represents a purchase order in the system
type Order struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	CropID      uint      `gorm:"not null" json:"crop_id"`
	Quantity    float64   `gorm:"not null" json:"quantity"`
	UnitPrice   float64   `gorm:"not null" json:"unit_price"`
	TotalPrice  float64   `gorm:"not null" json:"total_price"`
	Status      string    `gorm:"size:50;not null" json:"status"` // pending, confirmed, shipped, delivered, cancelled
	PaymentID   string    `gorm:"size:255" json:"payment_id"`
	Notes       string    `gorm:"size:1000" json:"notes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	Crop        Crop      `gorm:"foreignKey:CropID" json:"crop"`
}
