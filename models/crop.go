package models

import (
	"time"

	"github.com/lib/pq"
)

// Crop represents a crop that can be grown by users
type Crop struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	Name            string         `gorm:"size:255;not null" json:"name"` // indoor, outdoor
	FullImageURL    string         `gorm:"size:500" json:"full_image_url"`
	CroppedImageURL string         `gorm:"size:500" json:"cropped_image_url"`
	Description     string         `gorm:"size:1000" json:"description"`
	BasicNeeds      pq.StringArray `gorm:"type:text[]" json:"basic_needs"`
	Tags            pq.StringArray `gorm:"type:text[]" json:"tags"`
	Price           float64        `gorm:"not null" json:"price"`
	Rating          float64        `gorm:"default:0" json:"rating"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

// CropFavorite represents a user's favorite crop
type CropFavorite struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CropID    uint      `gorm:"not null" json:"crop_id"`
	Crop      Crop      `gorm:"foreignKey:CropID" json:"crop"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
