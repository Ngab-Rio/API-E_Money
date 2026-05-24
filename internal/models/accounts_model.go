package models

import "time"

type Account struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	Type      string    `gorm:"not null" json:"type"`
	Balance   float64   `gorm:"not null;default:0" json:"balance"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
