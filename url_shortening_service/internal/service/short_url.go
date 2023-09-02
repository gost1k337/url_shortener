package service

import (
	"context"
	"fmt"
	"github.com/gost1k337/url_shortener/url_shortening_service/config"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/hasher"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"math/rand"
	"time"
)

type ShortUrlService struct {
	repo   repository.ShortUrlRepo
	logger logging.Logger
	cfg    *config.Config
}

func NewShortUrlService(repo repository.ShortUrlRepo, logger logging.Logger, cfg *config.Config) *ShortUrlService {
	return &ShortUrlService{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

func (s *ShortUrlService) Create(ctx context.Context, userId int, originalUrl string, expireAt time.Duration) error {
	token := hasher.NewShortUrl(rand.Uint64())
	shortUrl := fmt.Sprintf("%s/%s", s.cfg.App.BaseURL, token)

	err := s.repo.Create(ctx, userId, originalUrl, shortUrl, expireAt)
	if err != nil {
		return err
	}

	return nil
}
