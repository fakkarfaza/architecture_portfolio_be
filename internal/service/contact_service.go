package service

import (
	"architecture_portfolio_api/internal/model"
	"architecture_portfolio_api/internal/repository"
	"context"
)

type ContactService struct {
	repository *repository.ContactRepository
}

func NewContactService(repository *repository.ContactRepository) *ContactService {
	return &ContactService{repository: repository}
}

func (s *ContactService) CreateContactMessage(
	ctx context.Context,
	message *model.ContactMessage,
) error {
	return s.repository.CreateContactMessage(ctx, message)
}
