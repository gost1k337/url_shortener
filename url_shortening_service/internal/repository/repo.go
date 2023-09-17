package repository

import (
	"context"
	"time"

	"github.com/gost1k337/url_shortener/url_shortening_service/internal/entity"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository/psql"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/postgres"
)

type ShortURL interface {
	Create(ctx context.Context, userId int, originalURL, shortURL string, expireAt time.Time) (int64, error)
	GetByID(ctx context.Context, id int64) (*entity.ShortURL, error)
	GetByShort(ctx context.Context, shortUrl string) (*entity.ShortURL, error)
}

type Repositories struct {
	ShortURL
}

func New(pg *postgres.Postgres, logger logging.Logger) *Repositories {
	return &Repositories{
		ShortURL: psql.NewShortURLRepo(pg, logger),
	}
}
