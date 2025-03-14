package models

import "gorm.io/gorm"

type Badges struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	Description    string `gorm:"type:text"`
	IconURL        string `gorm:"column:icon_url"`
	PointsRequired int    `gorm:"column:points_required"` // Points needed to unlock
}
