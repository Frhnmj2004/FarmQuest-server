package models

import "gorm.io/gorm"

type Alerts struct {
	gorm.Model
	UserID  int    `gorm:"not null"`
	Message string `gorm:"not null"`
	Type    string // info, warning, success
	IsRead  bool   `gorm:"default:false"`
	User    Users  `gorm:"foreignKey:UserID;references:ID"`
}
