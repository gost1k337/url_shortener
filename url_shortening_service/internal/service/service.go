package service

import (
	"context"
	"time"

	"github.com/gost1k337/url_shortener/url_shortening_service/config"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/entity"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
)

type ShortURL interface {
	Create(ctx context.Context, userId int, originalURL string, expireAt time.Time) (*entity.ShortURL, error)
	GetByShortToken(ctx context.Context, shortURLToken string) (*entity.ShortURL, error)
}

type Services struct {
	ShortUrl ShortURL
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps *ServicesDependencies, logger logging.Logger, cfg *config.Config) *Services {
	return &Services{
		ShortUrl: NewShortURLService(deps.Repos.ShortURL, logger, cfg),
	}
}
