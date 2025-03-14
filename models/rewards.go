package models

import "gorm.io/gorm"

type Rewards struct {
	gorm.Model
	UserID      int    `gorm:"not null"`
	Amount      int    `gorm:"not null"` // Points awarded
	Description string // e.g., Daily login, Harvest completed
	RewardType  string `gorm:"column:reward_type"` // points, badge, cash
	User        Users  `gorm:"foreignKey:UserID;references:ID"`
}
