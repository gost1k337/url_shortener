package service

import (
	"context"
	"github.com/gost1k337/url_shortener/url_shortening_service/config"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/entity"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/hasher"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"math/rand"
	"time"
)

type ShortUrlService struct {
	repo   repository.ShortUrl
	logger logging.Logger
	cfg    *config.Config
}

func NewShortUrlService(repo repository.ShortUrl, logger logging.Logger, cfg *config.Config) *ShortUrlService {
	return &ShortUrlService{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

func (s *ShortUrlService) Create(ctx context.Context, userId int, originalUrl string, expireAt time.Time) (*entity.ShortURL, error) {
	token := hasher.NewShortUrl(rand.Uint64())

	id, err := s.repo.Create(ctx, userId, originalUrl, token, expireAt)
	if err != nil {
		return nil, err
	}

	shortUrl, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return shortUrl, nil
}

func (s *ShortUrlService) GetByShortToken(ctx context.Context, shortUrlToken string) (*entity.ShortURL, error) {
	shortUrl, err := s.repo.GetByShort(ctx, shortUrlToken)
	if err != nil {
		return nil, err
	}

	return shortUrl, nil
}
