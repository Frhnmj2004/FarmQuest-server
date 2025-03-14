package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"type:text;not null"`
	AuthorID int    `gorm:"column:author_id"`
	Author   User  `gorm:"foreignKey:AuthorID;references:ID"`
}
