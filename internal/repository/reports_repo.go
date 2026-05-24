package repository

import (
	"context"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetOverview(ctx context.Context, userID int) (float64, float64, error)
	// GetMonthly(ctx context.Context, userID int) (float64, float64, error)
	// GetYearly(ctx context.Context, userID int) (float64, float64, error)
	// GetByCategory(ctx context.Context, userID int) (float64, float64, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (rr *reportRepository) GetOverview(ctx context.Context, userID int) (float64, float64, error) {
	var totalIncome float64
	var totalExpense float64

	err := rr.db.WithContext(ctx).
		Joins("JOIN accounts ON accounts.id = transactions.account_id").
		Where("accounts.user_id = ? AND transactions.type = ?", userID, "income").
		Select("COALESCE(SUM(transactions.amount), 0)").
		Scan(&totalIncome).Error
	if err != nil {
		return 0, 0, err
	}

	err = rr.db.WithContext(ctx).
		Joins("JOIN accounts ON accounts.id = transactions.account_id").
		Where("accounts.user_id = ? AND transactions.type = ?", userID, "expense").
		Select("COALESCE(SUM(transactions.amount), 0)").
		Scan(&totalExpense).Error
	if err != nil {
		return 0, 0, err
	}

	return totalIncome, totalExpense, nil
}
