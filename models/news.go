package models

import (
	"time"

	//"gorm.io/gorm"
)

type News struct {
	ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  int    `gorm:"column:author_id"`
	Author    User   `gorm:"foreignKey:AuthorID;references:ID"`
}
