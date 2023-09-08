package repository

import (
	"context"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/entity"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository/psql"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/postgres"
	"time"
)

type ShortUrl interface {
	Create(ctx context.Context, userId int, originalUrl, shortUrl string, expireAt time.Time) (int, error)
	GetById(ctx context.Context, id int) (*entity.ShortURL, error)
	GetByShort(ctx context.Context, shortUrl string) (*entity.ShortURL, error)
}

type Repositories struct {
	ShortUrl
}

func New(pg *postgres.Postgres, logger logging.Logger) *Repositories {
	return &Repositories{
		ShortUrl: psql.NewShortUrlRepo(pg, logger),
	}
}
