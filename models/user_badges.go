package models

import "time"

type UserBadges struct {
	UserID   int       `gorm:"primaryKey"`
	BadgeID  int       `gorm:"primaryKey"`
	EarnedAt time.Time `gorm:"default:now()"`
	User     User      `gorm:"foreignKey:UserID;references:ID"`
	Badge    Badge     `gorm:"foreignKey:BadgeID;references:ID"`
}
