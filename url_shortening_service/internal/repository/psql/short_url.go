package psql

import (
	"context"
	"fmt"
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

func (r *ShortUrlRepo) Create(ctx context.Context, userId int, originalUrl, shortUrl string, expireAt time.Duration) error {
	query := `INSERT INTO url_shorts (original_url, short_url, expire_at) VALUES ($1, $2, $3)`

	_, err := r.ExecContext(ctx, query, originalUrl, shortUrl, expireAt)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
