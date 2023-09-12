package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gost1k337/url_shortener/user_service/internal/entity"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
	"github.com/gost1k337/url_shortener/user_service/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
	logger logging.Logger
}

func NewUserRepo(pg *postgres.Postgres, logger logging.Logger) *UserRepo {
	return &UserRepo{
		pg,
		logger,
	}
}

func (r *UserRepo) Create(ctx context.Context, username, email, passwordHash string) (int, error) {
	query := `INSERT INTO users (username, email, passwordHash) VALUES ($1, $2, $3) RETURNING id`

	row := r.QueryRowContext(ctx, query, username, email, passwordHash)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("scan %w", err)
	}
	return id, nil
}

func (r *UserRepo) GetById(ctx context.Context, id int64) (*entity.User, error) {
	query := `SELECT * FROM users WHERE id=$1`

	row := r.QueryRowContext(ctx, query, id)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, nil
	}

	shortUrl := new(entity.User)

	if err := row.Scan(
		&shortUrl.Id,
		&shortUrl.Username,
		&shortUrl.Email,
		&shortUrl.PasswordHash,
		&shortUrl.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}
	return shortUrl, nil
}
