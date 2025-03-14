package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string  `gorm:"unique;not null"`
	Email    string  `gorm:"unique;not null"`
	Password string  `gorm:"not null"`
	Role     string  `gorm:"default:user"` // 'user' or 'admin'
	Points   int     `gorm:"default:0"`    // For gamification
	Balance  float64 `gorm:"default:0.00"` // Earnings from sales
}
