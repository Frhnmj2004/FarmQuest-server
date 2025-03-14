package models

import "time"

// Badge represents an achievement badge that users can earn
type Badge struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	Name           string    `gorm:"size:255;not null" json:"name"`
	Description    string    `gorm:"size:1000" json:"description"`
	IconURL        string    `gorm:"size:500" json:"icon_url"`
	PointsRequired int       `gorm:"not null" json:"points_required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// UserBadge represents a badge earned by a user
type UserBadge struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	BadgeID   uint      `gorm:"not null" json:"badge_id"`
	Badge     Badge     `gorm:"foreignKey:BadgeID" json:"badge"`
	EarnedAt  time.Time `gorm:"not null" json:"earned_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
