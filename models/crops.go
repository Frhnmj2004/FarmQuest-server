package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Crops struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	Type           string // e.g., vegetable, fruit, herb
	Popularity     int
	GrowthDuration int            `gorm:"column:growth_duration"` // Days to harvest
	Difficulty     string         // easy, medium, hard
	AvgYield       float64        `gorm:"column:avg_yield"` // Expected yield in kg
	Tags           pq.StringArray `gorm:"type:text[]"`
}
