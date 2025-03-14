package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model `gorm:"-"` // Ignore gorm.Model fields except UserID as primary key
	UserID     int        `gorm:"primaryKey;unique"`
	FullName   string     `gorm:"column:full_name"`
	Address    string     `gorm:"type:text"`
	Phone      string
	AvatarURL  string `gorm:"column:avatar_url"`
	Bio        string `gorm:"type:text"`
	User       Users  `gorm:"foreignKey:UserID;references:ID"`
}
