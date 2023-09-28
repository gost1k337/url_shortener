package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/gost1k337/url_shortener/url_shortening_service/config"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/entity"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/hasher"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
)

type ShortURLService struct {
	repo   repository.ShortURL
	logger logging.Logger
	cfg    *config.Config
}

func NewShortURLService(repo repository.ShortURL, logger logging.Logger, cfg *config.Config) *ShortURLService {
	return &ShortURLService{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

func (s *ShortURLService) Create(ctx context.Context, userId int, originalUrl string, expireAt time.Time) (
	*entity.ShortURL, error,
) {
	rMin := new(big.Int).SetInt64(1)
	rMax := new(big.Int).SetInt64(math.MaxInt64)

	r, err := rand.Int(rand.Reader, new(big.Int).Sub(rMax, rMin))
	if err != nil {
		return nil, fmt.Errorf("rand: %w", err)
	}

	token := hasher.NewShortURL(r.Uint64())

	id, err := s.repo.Create(ctx, userId, originalUrl, token, expireAt)
	if err != nil {
		return nil, fmt.Errorf("create urlshort: %w", err)
	}

	shortUrl, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get urlshort: %w", err)
	}

	return shortUrl, nil
}

func (s *ShortURLService) GetByShortToken(ctx context.Context, shortUrlToken string) (*entity.ShortURL, error) {
	shortUrl, err := s.repo.GetByShort(ctx, shortUrlToken)
	if err != nil {
		return nil, fmt.Errorf("get urlshort: %w", err)
	}

	return shortUrl, nil
}
