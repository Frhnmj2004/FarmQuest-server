package models

import (
	"time"
	"gorm.io/gorm"
)

type Farms struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ImageURL    string         `json:"image_url"`
	UserID      int           `gorm:"not null" json:"user_id"`
	CropID      int           `gorm:"not null" json:"crop_id"`
	Description string         `gorm:"size:1000" json:"description"`
	Location    string         `gorm:"size:255" json:"location"`
	Name        string         `gorm:"size:255" json:"name"`
	Status      string         `json:"status"` // planted, growing, harvesting, completed
	PlantedAt   time.Time      `json:"planted_at"`
	GrowingAt   time.Time      `json:"growing_at,omitempty"`
	HarvestAt   time.Time      `json:"harvest_at,omitempty"`
	Health      int           `gorm:"default:100" json:"health"`
	Area        float64       `gorm:"default:0" json:"area"`
	User        User          `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Crop        Crop          `gorm:"foreignKey:CropID;references:ID" json:"crop"`
}
