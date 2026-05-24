package services

import (
	"context"
	"e-money/internal/models"
	"e-money/internal/repository"
)

type AccountService interface {
	GetAccountByUserID(ctx context.Context, UserID int) ([]models.Account, error)
	GetSummary(ctx context.Context, UserID int) (float64, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{accountRepo: accountRepo}
}

func (s *accountService) GetAccountByUserID(ctx context.Context, UserID int) ([]models.Account, error) {
	return s.accountRepo.GetAccountByUserID(ctx, UserID)
}

func (s *accountService) GetSummary(ctx context.Context, UserID int) (float64, error) {
	return s.accountRepo.GetSummary(ctx, UserID)
}
