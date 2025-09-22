package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint   	   `json:"user_id" gorm:"not null"`
	User        User   	   `json:"-" gorm:"foreignKey:UserID"`
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description"`
	Completed   bool           `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

