package models

import (
	"time"

	//"gorm.io/gorm"
)

type Farms struct {
	ID        	uint `gorm:"primarykey"`
    CreatedAt 	time.Time
    UpdatedAt 	time.Time
    //DeletedAt 	time.Time `gorm:"index"`
	UserID      int `gorm:"not null"`
	CropID      int `gorm:"not null"`
	Name        string
	Status      string // planted, growing, harvesting, completed
	PlantedAt   time.Time
	HarvestedAt time.Time
	Yield       float64 `gorm:"default:0"` // Actual yield in kg
	User        User   `gorm:"foreignKey:UserID;references:ID"`
	Crop        Crop   `gorm:"foreignKey:CropID;references:ID"`
}
