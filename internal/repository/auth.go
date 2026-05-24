package repository

import (
	"context"
	"e-money/internal/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(ctx context.Context, auth *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id int) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (au *authRepository) Create(ctx context.Context, auth *models.User) error {
	return au.db.WithContext(ctx).Create(auth).Error
}

func (au *authRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := au.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (au *authRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	err := au.db.WithContext(ctx).First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
