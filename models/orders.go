package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	UserID   int     `gorm:"not null"` // Seller
	BuyerID  int     `gorm:"not null"`
	FarmID   int     `gorm:"not null"`
	Quantity float64 `gorm:"not null"` // kg ordered
	Price    float64 `gorm:"not null"`
	Status   string  // pending, shipped, delivered, cancelled
	Seller   Users   `gorm:"foreignKey:UserID;references:ID"`
	Buyer    Users   `gorm:"foreignKey:BuyerID;references:ID"`
	Farm     Farms   `gorm:"foreignKey:FarmID;references:ID"`
}
