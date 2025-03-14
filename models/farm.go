package models

import (
	"time"

	"gorm.io/gorm"
)

type Farms struct {
	gorm.Model
	UserID      int `gorm:"not null"`
	CropID      int `gorm:"not null"`
	Name        string
	Status      string // planted, growing, harvesting, completed
	PlantedAt   time.Time
	HarvestedAt time.Time
	Yield       float64 `gorm:"default:0"` // Actual yield in kg
	User        Users   `gorm:"foreignKey:UserID;references:ID"`
	Crop        Crops   `gorm:"foreignKey:CropID;references:ID"`
}
