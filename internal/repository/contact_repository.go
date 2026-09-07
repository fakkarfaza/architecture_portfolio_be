package repository

import (
	"architecture_portfolio_api/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ContactRepository struct {
	db *pgxpool.Pool
}

func NewContactRepository(db *pgxpool.Pool) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) CreateContactMessage(
	ctx context.Context,
	message *model.ContactMessage,
) error {
	query := `
INSERT INTO contact_messages (
                             name,
                             email,
                             subject,
                             message
)
Values ($1, $2, $3, $4)
`

	_, err := r.db.Exec(ctx, query, message.Name, message.Email, message.Subject, message.Message)

	return err
}
