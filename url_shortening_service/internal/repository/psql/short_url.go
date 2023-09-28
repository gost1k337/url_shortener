package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gost1k337/url_shortener/url_shortening_service/internal/entity"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/postgres"
)

type ShortURLRepo struct {
	*postgres.Postgres
	logger logging.Logger
}

func NewShortURLRepo(pg *postgres.Postgres, logger logging.Logger) *ShortURLRepo {
	return &ShortURLRepo{
		pg,
		logger,
	}
}

func (r *ShortURLRepo) Create(ctx context.Context, userId int, originalURL, shortURL string, expireAt time.Time) (
	int64, error,
) {
	query := `INSERT INTO url_shorts (original_url, short_url, visits, expire_at) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64

	err := r.QueryRowContext(ctx, query, originalURL, shortURL, 0, expireAt).Scan(&id) //nolint:execinquery
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}

	return id, nil
}

func (r *ShortURLRepo) GetByID(ctx context.Context, id int64) (*entity.ShortURL, error) {
	query := `SELECT * FROM url_shorts WHERE id=$1`

	row := r.QueryRowContext(ctx, query, id)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, fmt.Errorf("no rows")
	}

	shortUrl := new(entity.ShortURL)

	if err := row.Scan(
		&shortUrl.Id,
		&shortUrl.OriginalURL,
		&shortUrl.ShortURL,
		&shortUrl.Visits,
		&shortUrl.ExpireAt,
		&shortUrl.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return shortUrl, nil
}

func (r *ShortURLRepo) GetByShort(ctx context.Context, shortUrlToken string) (*entity.ShortURL, error) {
	query := `SELECT * FROM url_shorts WHERE short_url=$1`

	row := r.QueryRowContext(ctx, query, shortUrlToken)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, fmt.Errorf("no rows")
	}

	shortUrl := new(entity.ShortURL)
	if err := row.Scan(
		&shortUrl.Id,
		&shortUrl.OriginalURL,
		&shortUrl.ShortURL,
		&shortUrl.Visits,
		&shortUrl.ExpireAt,
		&shortUrl.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return shortUrl, nil
}
