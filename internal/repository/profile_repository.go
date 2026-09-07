package repository

import (
	"architecture_portfolio_api/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) GetProfile(
	ctx context.Context,
	locale string,
) (*model.Profile, error) {

	query := `
		SELECT
			p.id,
			p.name,
			p.profile_image_url,
			pt.title,
			pt.bio,
			pt.description,
			p.email,
			p.location,
			p.instagram_url,
			p.linkedin_url,
			p.behance_url
		FROM profiles p
		LEFT JOIN profile_translations pt
			ON pt.profile_id = p.id
			AND pt.locale = $1
		ORDER BY p.id
		LIMIT 1
	`

	var profile model.Profile

	err := r.db.QueryRow(
		ctx,
		query,
		locale,
	).Scan(
		&profile.ID,
		&profile.Name,
		&profile.ProfileImageURL,
		&profile.Title,
		&profile.Bio,
		&profile.Description,
		&profile.Email,
		&profile.Location,
		&profile.InstagramURL,
		&profile.LinkedInURL,
		&profile.BehanceURL,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}
