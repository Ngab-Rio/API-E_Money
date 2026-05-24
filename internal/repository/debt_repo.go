package repository

import (
	"context"
	"e-money/internal/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type DebtRepository interface {
	CreateDebt(ctx context.Context, debt *models.Debt) error
	GetDebtByID(ctx context.Context, userID int, id int) (*models.Debt, error)
	GetDebtByUserID(ctx context.Context, userID int, filter models.FilterDebt) ([]models.Debt, error)
	UpdateDebt(ctx context.Context, debt *models.Debt) error
	UpdateDebtPaid(ctx context.Context, id int, userID int, accountType string) error
	DeleteDebt(ctx context.Context, id int, userID int) error
}

type debtRepository struct {
	db *gorm.DB
}

func NewDebtRepository(db *gorm.DB) DebtRepository {
	return &debtRepository{db: db}
}

func (dt *debtRepository) CreateDebt(ctx context.Context, debt *models.Debt) error {
	return dt.db.WithContext(ctx).Create(debt).Error
}

func (dt *debtRepository) GetDebtByID(ctx context.Context, userID int, id int) (*models.Debt, error) {
	var debt models.Debt
	err := dt.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&debt).Error
	if err != nil {
		return nil, err
	}
	return &debt, nil
}

func (dt *debtRepository) GetDebtByUserID(ctx context.Context, userID int, filter models.FilterDebt) ([]models.Debt, error) {
	var debts []models.Debt
	db := dt.db.WithContext(ctx).Where("user_id = ?", userID)

	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	if filter.Type != "" {
		db = db.Where("type = ?", filter.Type)
	}

	err := db.Order("created_at DESC").Find(&debts).Error
	if err != nil {
		return nil, err
	}
	return debts, nil
}

func (dt *debtRepository) UpdateDebt(ctx context.Context, debt *models.Debt) error {
	return dt.db.WithContext(ctx).Where("id = ? AND user_id = ?", debt.ID, debt.UserID).Updates(debt).Error
}

func (dt *debtRepository) UpdateDebtPaid(ctx context.Context, id int, userID int, accountType string) error {
	return dt.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var debt models.Debt
		if err := tx.Where("id = ? AND user_id = ?", id, userID).First(&debt).Error; err != nil {
			return err
		}

		if debt.Status == "paid" {
			return errors.New("debt is already marked as paid")
		}

		var account models.Account
		err := tx.Where("user_id = ? AND type = ?", userID, accountType).First(&account).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("target account (cash/bank) not found for this user")
			}
			return err
		}

		var operator string
		var transactionType string
		var txDesc string

		switch debt.Type {
		case "debt":
			operator = "-"
			transactionType = "expense"
			txDesc = "Pelunasan hutang ke " + debt.PersonName
		case "loan":
			operator = "+"
			transactionType = "income"
			txDesc = "Pemberian pinjaman dari " + debt.PersonName
		default:
			return errors.New("invalid debt type")
		}

		result := tx.Model(&models.Account{}).
			Where("id = ?", account.ID).
			Update("balance", gorm.Expr("balance "+operator+" ?", debt.Amount))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("target account not found")
		}

		err = tx.Exec(`INSERT INTO transactions (account_id, amount, category, type, note) VALUES (?, ?, ?, ?, ?)`,
			account.ID, debt.Amount, debt.Type, transactionType, txDesc).Error
		if err != nil {
			return err
		}
		now := time.Now()
		updates := map[string]interface{}{
			"status":       "paid",
			"account_type": accountType,
			"paid_at":      &now,
		}

		return tx.Model(&debt).Updates(updates).Error
	})
}

func (dt *debtRepository) DeleteDebt(ctx context.Context, debtID int, userID int) error {
	res := dt.db.WithContext(ctx).Where("id = ? AND user_id = ?", debtID, userID).Delete(&models.Debt{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
