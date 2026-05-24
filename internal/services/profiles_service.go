package services

import (
	"context"
	"e-money/internal/models"
	"e-money/internal/repository"
)

type ProfileService interface {
	GetProfileByUserID(ctx context.Context, UserID int) (*models.Profile, error)
}

type profileService struct {
	profileRepo repository.ProfileRepository
}

func NewProfileService(profileRepo repository.ProfileRepository) ProfileService {
	return &profileService{profileRepo: profileRepo}
}

func (s *profileService) GetProfileByUserID(ctx context.Context, UserID int) (*models.Profile, error) {
	return s.profileRepo.GetProfileByUserID(ctx, UserID)
}
