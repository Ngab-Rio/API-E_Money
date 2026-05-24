package services

import (
	"context"
	"e-money/internal/dto"
	"e-money/internal/models"
	"e-money/internal/repository"
	"e-money/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
}

type authService struct {
	authRepo    repository.AuthRepository
	accountRepo repository.AccountRepository
	jwtManager  utils.JWTManager
}

func NewAuthService(authRepo repository.AuthRepository, accountRepo repository.AccountRepository, jwtManager utils.JWTManager) AuthService {
	return &authService{authRepo: authRepo, accountRepo: accountRepo, jwtManager: jwtManager}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existingUser, err := s.authRepo.FindByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err // error beneran (DB down, dll)
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.authRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := s.jwtManager.Generate(user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	if err := s.accountRepo.CreateAccount(ctx, &models.Account{
		UserID:  user.ID,
		Type:    "cash",
		Balance: 0,
	}); err != nil {
		return nil, err
	}

	if err := s.accountRepo.CreateAccount(ctx, &models.Account{
		UserID:  user.ID,
		Type:    "bank",
		Balance: 0,
	}); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Token:  token,
	}, nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.authRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := s.jwtManager.Generate(user.ID, user.Name)

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Token:  token,
	}, nil
}
