package service

import (
	"architecture_portfolio_api/internal/model"
	"architecture_portfolio_api/internal/repository"
	"context"
)

type ProjectService struct {
	repository *repository.ProjectRepository
}

func NewProjectService(repository *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repository}
}

func (s *ProjectService) GetProjects(
	ctx context.Context,
	locale string,
) ([]model.Project, error) {
	return s.repository.GetProjects(ctx, locale)
}

func (s *ProjectService) GetProjectBySlug(
	ctx context.Context,
	slug string,
	locale string,
) (*model.ProjectDetail, error) {
	return s.repository.GetProjectBySlug(
		ctx,
		slug,
		locale,
	)
}
