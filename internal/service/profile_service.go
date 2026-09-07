package service

import (
	"architecture_portfolio_api/internal/model"
	"architecture_portfolio_api/internal/repository"
	"context"
)

type ProfileService struct {
	repository *repository.ProfileRepository
}

func NewProfileService(repository *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repository: repository}
}

func (s *ProfileService) GetProfile(
	ctx context.Context,
	locale string,
) (*model.Profile, error) {
	return s.repository.GetProfile(ctx, locale)
}
