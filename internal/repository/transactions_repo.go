package repository

import (
	"context"
	"e-money/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	GetTransactionByID(ctx context.Context, userID int, transactionID int) (*models.Transaction, error)
	GetAllTransactionsByUserID(ctx context.Context, UserID int, filter models.FilterTransaction) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (tr *transactionRepository) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	return tr.db.WithContext(ctx).Create(transaction).Error
}

func (tr *transactionRepository) GetTransactionByID(ctx context.Context, userID int, transactionID int) (*models.Transaction, error) {
	var transaction models.Transaction
	err := tr.db.WithContext(ctx).Where("id = ? AND user_id = ?", transactionID, userID).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (tr *transactionRepository) GetAllTransactionsByUserID(ctx context.Context, UserID int, filter models.FilterTransaction) ([]models.Transaction, error) {
	var transactions []models.Transaction

	db := tr.db.WithContext(ctx).
		Joins("JOIN accounts ON accounts.id = transactions.account_id").
		Where("accounts.user_id = ?", UserID)

	if filter.Type != "" {
		db = db.Where("transactions.type = ?", filter.Type)
	}

	if filter.AccountType != "" {
		db = db.Where("accounts.type = ?", filter.AccountType)
	}

	err := db.Order("transactions.created_at DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
