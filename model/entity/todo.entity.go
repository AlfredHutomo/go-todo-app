package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `json:"deleted_at"  gorm:"index"`
}
