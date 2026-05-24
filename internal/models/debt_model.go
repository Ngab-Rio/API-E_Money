package models

import "time"

type Debt struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int        `gorm:"not null" json:"user_id"`
	AccountType *string    `gorm:"not null" json:"account_type"`
	PersonName  string     `gorm:"not null" json:"person_name"`
	Type        string     `gorm:"not null" json:"type"`
	Amount      float64    `gorm:"not null;default:0" json:"amount"`
	Status      string     `gorm:"not null;default:'unpaid'" json:"status"`
	Note        string     `gorm:"not null" json:"note"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	PaidAt      *time.Time `gorm:"null" json:"paid_at"`
}

type FilterDebt struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}
