package models

import "time"

type Transaction struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	AccountID int       `gorm:"not null" json:"account_id"`
	Type      string    `gorm:"not null" json:"type"`
	Amount    float64   `gorm:"not null;default:0" json:"amount"`
	Category  string    `gorm:"not null" json:"category"`
	Note      string    `gorm:"not null" json:"note"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type FilterTransaction struct {
	Type        string `json:"type"`
	Category    string `json:"category"`
	AccountType string `json:"account_type"`
}
