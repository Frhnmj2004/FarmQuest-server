package models

import (
	"time"

	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	FarmID       int    `gorm:"not null"`
	TaskType     string `gorm:"column:task_type"` // water, fertilize, check, harvest
	PointsReward int    `gorm:"default:10"`
	Status       string `gorm:"default:pending"` // pending, completed
	DueAt        time.Time
	CompletedAt  time.Time
	Farm         Farms `gorm:"foreignKey:FarmID;references:ID"`
}
