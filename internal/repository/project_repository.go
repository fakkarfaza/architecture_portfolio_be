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

func (r *ProjectRepository) GetProjectBySlug(
	ctx context.Context,
	slug string,
	locale string,
) (*model.ProjectDetail, error) {

	query := `
		SELECT
			p.id,
			p.slug,
			pt.title,
			pt.category,
			pt.location,
			p.year,
			p.thumbnail_image_url,
			p.hero_image_url,
			p.concept_image_url,
			pt.brief,
			pt.concept_title,
			pt.concept
		FROM projects p
		LEFT JOIN project_translations pt
			ON pt.project_id = p.id
			AND pt.locale = $2
		WHERE p.slug = $1
		LIMIT 1
	`

	var project model.ProjectDetail

	err := r.db.QueryRow(
		ctx,
		query,
		slug,
		locale,
	).Scan(
		&project.ID,
		&project.Slug,
		&project.Title,
		&project.Category,
		&project.Location,
		&project.Year,
		&project.ThumbnailImageURL,
		&project.HeroImageURL,
		&project.ConceptImageURL,
		&project.Brief,
		&project.ConceptTitle,
		&project.Concept,
	)

	if err != nil {
		return nil, err
	}

	galleryQuery := `
		SELECT
			id,
			image_url,
			image_type,
			sort_order
		FROM project_images
		WHERE project_id = $1
		ORDER BY sort_order ASC, id ASC
	`

	rows, err := r.db.Query(
		ctx,
		galleryQuery,
		project.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	project.Gallery = make([]model.ProjectImage, 0)

	for rows.Next() {
		var image model.ProjectImage

		err := rows.Scan(
			&image.ID,
			&image.ImageURL,
			&image.ImageType,
			&image.SortOrder,
		)
		if err != nil {
			return nil, err
		}

		project.Gallery = append(project.Gallery, image)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &project, nil
}
