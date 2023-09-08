package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/entity"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/postgres"
	"time"
)

type ShortUrlRepo struct {
	*postgres.Postgres
	logger logging.Logger
}

func NewShortUrlRepo(pg *postgres.Postgres, logger logging.Logger) *ShortUrlRepo {
	return &ShortUrlRepo{
		pg,
		logger,
	}
}

func (r *ShortUrlRepo) Create(ctx context.Context, userId int, originalUrl, shortUrl string, expireAt time.Time) (int, error) {
	query := `INSERT INTO url_shorts (original_url, short_url, visits, expire_at) VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.QueryRowContext(ctx, query, originalUrl, shortUrl, 0, expireAt)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("scan %w", err)
	}
	return id, nil
}

func (r *ShortUrlRepo) GetById(ctx context.Context, id int) (*entity.ShortURL, error) {
	query := `SELECT * FROM url_shorts WHERE id=$1`

	row := r.QueryRowContext(ctx, query, id)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, nil
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

func (r *ShortUrlRepo) GetByShort(ctx context.Context, shortUrlToken string) (*entity.ShortURL, error) {
	query := `SELECT * FROM url_shorts WHERE short_url=$1`

	row := r.QueryRowContext(ctx, query, shortUrlToken)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, nil
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
