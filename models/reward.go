package models

import (
	"time"
)

// Reward represents a user reward in the system
type Reward struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	Type        string    `gorm:"size:50;not null" json:"type"` // achievement, daily, special
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"size:1000" json:"description"`
	Points      int       `gorm:"not null" json:"points"`
	Claimed     bool      `gorm:"default:false" json:"claimed"`
	ClaimedAt   time.Time `json:"claimed_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
}
