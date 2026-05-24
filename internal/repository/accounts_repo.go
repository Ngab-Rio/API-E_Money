package repository

import (
	"context"
	"e-money/internal/models"

	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountByUserID(ctx context.Context, UserID int) ([]models.Account, error)
	GetSummary(ctx context.Context, UserID int) (float64, error)
	CreateAccount(ctx context.Context, account *models.Account) error
	FindByUserIDAndType(ctx context.Context, userID int, Type string) (*models.Account, error)
	Update(ctx context.Context, account *models.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (au *accountRepository) CreateAccount(ctx context.Context, account *models.Account) error {
	return au.db.WithContext(ctx).Create(account).Error
}

func (au *accountRepository) GetAccountByUserID(ctx context.Context, UserID int) ([]models.Account, error) {
	var account []models.Account
	err := au.db.WithContext(ctx).Where("user_id = ?", UserID).Find(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (au *accountRepository) GetSummary(ctx context.Context, UserID int) (float64, error) {
	var total float64
	err := au.db.WithContext(ctx).Table("accounts").Where("user_id = ?", UserID).Select("COALESCE(SUM(balance), 0) as balance").Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (au *accountRepository) FindByUserIDAndType(ctx context.Context, UserID int, Type string) (*models.Account, error) {
	var account models.Account
	err := au.db.WithContext(ctx).Where("user_id = ? AND type = ?", UserID, Type).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (au *accountRepository) Update(ctx context.Context, account *models.Account) error {
	return au.db.WithContext(ctx).Model(&models.Account{}).Where("id = ?", account.ID).Update("balance", account.Balance).Error
}
