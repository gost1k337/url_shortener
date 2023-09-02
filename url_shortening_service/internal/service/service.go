package service

import (
	"context"
	"github.com/gost1k337/url_shortener/url_shortening_service/config"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"time"
)

type ShortUrl interface {
	Create(ctx context.Context, userId int, originalUrl string, expireAt time.Duration) error
}

type Services struct {
	ShortUrl ShortUrl
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps *ServicesDependencies, logger logging.Logger, cfg *config.Config) *Services {
	return &Services{
		ShortUrl: NewShortUrlService(deps.Repos.ShortUrlRepo, logger, cfg),
	}
}
