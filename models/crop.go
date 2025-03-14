package models

import "time"

// Crop represents a crop that can be grown by users
type Crop struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Name         string    `gorm:"size:255;not null" json:"name"`
	Type         string    `gorm:"size:100;not null" json:"type"`
	Category     string    `gorm:"size:50;not null" json:"category"` // indoor, outdoor
	ImageURL     string    `gorm:"size:500" json:"image_url"`
	Description  string    `gorm:"size:1000" json:"description"`
	WaterNeed    int       `gorm:"not null" json:"water_need"`    // 1-5 scale
	SunlightNeed int       `gorm:"not null" json:"sunlight_need"` // 1-5 scale
	Price        float64   `gorm:"not null" json:"price"`
	IsPopular    bool      `gorm:"default:false" json:"is_popular"`
	Popularity   int       `gorm:"default:0" json:"popularity"`
	Rating       float64   `gorm:"default:0" json:"rating"`
	ReviewCount  int       `gorm:"default:0" json:"review_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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
