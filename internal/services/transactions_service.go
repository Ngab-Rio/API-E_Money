package services

import (
	"context"
	"e-money/internal/dto"
	"e-money/internal/models"
	"e-money/internal/repository"
	"errors"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, UserID int, req dto.CreateTransactionRequest) error
	GetTransactionByID(ctx context.Context, userID int, transactionID int) (*models.Transaction, error)
	GetAllTransactionsByUserID(ctx context.Context, UserID int, filter models.FilterTransaction) ([]models.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, accountRepo repository.AccountRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo, accountRepo: accountRepo}
}

func (s *transactionService) CreateTransaction(ctx context.Context, UserID int, req dto.CreateTransactionRequest) error {
	account, err := s.accountRepo.FindByUserIDAndType(ctx, UserID, req.AccountType)
	if err != nil {
		return err
	}

	if account == nil {
		return errors.New("Account not found")
	}

	if req.Type == "expense" && account.Balance < req.Amount {
		return errors.New("Insufficient balance")
	}

	newBalance := account.Balance
	if req.Type == "expense" {
		newBalance -= req.Amount
	} else {
		newBalance += req.Amount
	}

	err = s.transactionRepo.CreateTransaction(ctx, &models.Transaction{
		AccountID: account.ID,
		UserID:    UserID,
		Type:      req.Type,
		Amount:    req.Amount,
		Category:  req.Category,
		Note:      req.Note,
	})

	return s.accountRepo.Update(ctx, &models.Account{
		ID:      account.ID,
		Balance: newBalance,
	})
}

func (s *transactionService) GetTransactionByID(ctx context.Context, userID int, transactionID int) (*models.Transaction, error) {
	return s.transactionRepo.GetTransactionByID(ctx, userID, transactionID)
}

func (s *transactionService) GetAllTransactionsByUserID(ctx context.Context, UserID int, filter models.FilterTransaction) ([]models.Transaction, error) {
	return s.transactionRepo.GetAllTransactionsByUserID(ctx, UserID, filter)
}
