package repository

import (
	"architecture_portfolio_api/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) GetProjects(
	ctx context.Context,
	locale string,
) ([]model.Project, error) {

	query := `
		SELECT
			p.id,
			p.slug,
			pt.title,
			pt.category,
			pt.location,
			p.year,
			p.thumbnail_image_url,
			p.hero_image_url
		FROM projects p
		LEFT JOIN project_translations pt
			ON pt.project_id = p.id
			AND pt.locale = $1
		ORDER BY p.year DESC NULLS LAST, p.id DESC
`

	rows, err := r.db.Query(ctx, query, locale)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project

	for rows.Next() {
		var project model.Project

		err := rows.Scan(
			&project.ID,
			&project.Slug,
			&project.Title,
			&project.Category,
			&project.Location,
			&project.Year,
			&project.ThumbnailImageURL,
			&project.HeroImageURL,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
