package models

import "time"

// Task represents a farming task that users can complete
type Task struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"size:1000" json:"description"`
	Type        string    `gorm:"size:50;not null" json:"type"` // daily, weekly, achievement
	Points      int       `gorm:"not null" json:"points"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	CompletedAt time.Time `json:"completed_at"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
