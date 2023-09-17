package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	us "github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type URLShortService struct {
	c      us.UrlShortsClient
	logger logging.Logger
}

func NewURLShortService(uc us.UrlShortsClient, logger logging.Logger) *URLShortService {
	return &URLShortService{
		c:      uc,
		logger: logger,
	}
}

func (s *URLShortService) Create(ctx context.Context, originalURL string, expireAt time.Duration) (
	*CreateURLShortResp, error,
) {
	createURLShortReq := &us.CreateUrlShortRequest{
		OriginalUrl: originalURL,
		ExpireAt:    timestamppb.New(time.Now().Add(expireAt)),
	}

	urlshort, err := s.c.Create(ctx, createURLShortReq)
	if err != nil {
		return nil, fmt.Errorf("url-short create: %w", err)
	}

	res := &CreateURLShortResp{
		ID:          urlshort.Id,
		OriginalURL: urlshort.OriginalUrl,
		ShortURL:    urlshort.ShortUrl,
		Visits:      urlshort.Visits,
		ExpireAt:    urlshort.ExpireAt.AsTime(),
		CreatedAt:   urlshort.CreatedAt.AsTime(),
	}

	return res, nil
}

func (s *URLShortService) Get(ctx context.Context, token string) (*GetURLShortResp, error) {
	getURLShortReq := &us.GetUrlShortRequest{
		Token: token,
	}

	urlshort, err := s.c.Get(ctx, getURLShortReq)
	if err != nil {
		return nil, fmt.Errorf("url-short get: %w", err)
	}

	res := &GetURLShortResp{
		ID:          urlshort.Id,
		OriginalURL: urlshort.OriginalUrl,
		ShortURL:    urlshort.ShortUrl,
		Visits:      urlshort.Visits,
		ExpireAt:    urlshort.ExpireAt.AsTime(),
		CreatedAt:   urlshort.CreatedAt.AsTime(),
	}

	return res, nil
}
