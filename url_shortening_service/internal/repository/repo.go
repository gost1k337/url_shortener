package repository

import (
	"context"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository/psql"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/postgres"
	"time"
)

type ShortUrlRepo interface {
	Create(ctx context.Context, userId int, originalUrl, shortUrl string, expireAt time.Duration) error
}

type Repositories struct {
	ShortUrlRepo
}

func New(pg *postgres.Postgres, logger logging.Logger) *Repositories {
	return &Repositories{
		ShortUrlRepo: psql.NewShortUrlRepo(pg, logger),
	}
}
