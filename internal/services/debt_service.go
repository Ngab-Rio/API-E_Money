package services

import (
	"context"
	"e-money/internal/dto"
	"e-money/internal/models"
	"e-money/internal/repository"
	"errors"
)

type DebtService interface {
	CreateDebt(ctx context.Context, userID int, req dto.CreateDebtRequest) error
	GetDebtByID(ctx context.Context, userID int, id int) (*models.Debt, error)
	GetDebtByUserID(ctx context.Context, userID int, filter models.FilterDebt) ([]models.Debt, error)
	UpdateDebt(ctx context.Context, userID int, debtID int, req dto.UpdateDebtRequest) error
	DeleteDebt(ctx context.Context, id int, userID int) error
}

type debtService struct {
	debtRepo repository.DebtRepository
}

func NewDebtService(debtRepo repository.DebtRepository) DebtService {
	return &debtService{debtRepo: debtRepo}
}

func (s *debtService) CreateDebt(ctx context.Context, userID int, req dto.CreateDebtRequest) error {
	if req.PersonName == "" {
		return errors.New("person name is required")
	}
	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	if req.Type != "debt" && req.Type != "loan" {
		return errors.New("invalid type")
	}
	if req.AccountType != "cash" && req.AccountType != "bank" {
		return errors.New("invalid account type, must be 'cash' or 'bank'")
	}

	debt := &models.Debt{
		UserID:      userID,
		AccountType: &req.AccountType,
		PersonName:  req.PersonName,
		Amount:      req.Amount,
		Type:        req.Type,
		Status:      "unpaid",
		Note:        req.Note,
	}

	if err := s.debtRepo.CreateDebt(ctx, debt); err != nil {
		return nil
	}

	return nil
}

func (s *debtService) GetDebtByID(ctx context.Context, userID int, id int) (*models.Debt, error) {
	return s.debtRepo.GetDebtByID(ctx, userID, id)
}

func (s *debtService) GetDebtByUserID(ctx context.Context, userID int, filter models.FilterDebt) ([]models.Debt, error) {
	return s.debtRepo.GetDebtByUserID(ctx, userID, filter)
}

func (s *debtService) UpdateDebt(ctx context.Context, userID int, debtID int, req dto.UpdateDebtRequest) error {
	debt, err := s.debtRepo.GetDebtByID(ctx, userID, debtID)
	if err != nil {
		return err
	}
	if req.Status != nil && *req.Status == "paid" {
		if debt.AccountType == nil || *debt.AccountType == "" {
			return errors.New("account type is required")
		}
		return s.debtRepo.UpdateDebtPaid(ctx, debtID, userID, *debt.AccountType)
	}
	if req.PersonName != nil {
		if *req.PersonName == "" {
			return errors.New("person name is required")
		}
		debt.PersonName = *req.PersonName
	}
	if req.Amount != nil {
		if *req.Amount <= 0 {
			return errors.New("amount must be greater than 0")
		}
		debt.Amount = *req.Amount
	}
	if req.Type != nil {
		if *req.Type != "debt" && *req.Type != "loan" {
			return errors.New("invalid type")
		}
		debt.Type = *req.Type
	}
	if req.Note != nil {
		debt.Note = *req.Note
	}

	if err := s.debtRepo.UpdateDebt(ctx, debt); err != nil {
		return err
	}

	return nil
}

func (s *debtService) DeleteDebt(ctx context.Context, id int, userID int) error {
	return s.debtRepo.DeleteDebt(ctx, id, userID)
}
