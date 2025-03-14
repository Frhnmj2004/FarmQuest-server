package models

import (
	"time"
)

// Profile represents additional user information
type Profile struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null;unique" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	FullName  string    `gorm:"size:255" json:"full_name"`
	Address   string    `gorm:"size:500" json:"address"`
	Phone     string    `gorm:"size:50" json:"phone"`
	AvatarURL string    `gorm:"size:500" json:"avatar_url"`
	Bio       string    `gorm:"size:1000" json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
