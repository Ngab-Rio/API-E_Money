package models

import "time"

type Profile struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
