package repository

import (
	"context"
	"e-money/internal/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfileByUserID(ctx context.Context, UserID int) (*models.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (pr *profileRepository) GetProfileByUserID(ctx context.Context, UserID int) (*models.Profile, error) {
	var profile models.Profile
	err := pr.db.WithContext(ctx).Table("users").Where("id = ?", UserID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
