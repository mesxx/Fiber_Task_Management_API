package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID      uint           `gorm:"not null;" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	Status      string         `gorm:"not null;" json:"status"`
	Title       string         `gorm:"not null;" json:"title"`
	Description sql.NullString `gorm:"default:NULL;" json:"description"`
	Image       sql.NullString `gorm:"default:NULL;" json:"image"`
}
