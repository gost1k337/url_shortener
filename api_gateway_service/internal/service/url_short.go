package service

import (
	"context"
	"fmt"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	us "github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type UrlShortService struct {
	c      us.UrlShortsClient
	logger logging.Logger
}

func NewUrlShortService(uc us.UrlShortsClient, logger logging.Logger) *UrlShortService {
	return &UrlShortService{
		c:      uc,
		logger: logger,
	}
}

func (s *UrlShortService) Create(ctx context.Context, originalUrl string, expireAt time.Duration) (*CreateUrlShortResp, error) {
	createUrlShortReq := &us.CreateUrlShortRequest{
		OriginalUrl: originalUrl,
		ExpireAt:    timestamppb.New(time.Now().Add(expireAt)),
	}
	u, err := s.c.Create(ctx, createUrlShortReq)
	if err != nil {
		return nil, fmt.Errorf("url-short create: %w", err)
	}
	res := &CreateUrlShortResp{
		Id:          u.Id,
		OriginalUrl: u.OriginalUrl,
		ShortUrl:    u.ShortUrl,
		Visits:      u.Visits,
		ExpireAt:    u.ExpireAt.AsTime(),
		CreatedAt:   u.CreatedAt.AsTime(),
	}
	return res, nil
}

func (s *UrlShortService) Get(ctx context.Context, token string) (*GetUrlShortResp, error) {
	getUrlShortReq := &us.GetUrlShortRequest{
		Token: token,
	}

	us, err := s.c.Get(ctx, getUrlShortReq)
	if err != nil {
		return nil, fmt.Errorf("url-short get: %w", err)
	}
	res := &GetUrlShortResp{
		Id:          us.Id,
		OriginalUrl: us.OriginalUrl,
		ShortUrl:    us.ShortUrl,
		Visits:      us.Visits,
		ExpireAt:    us.ExpireAt.AsTime(),
		CreatedAt:   us.CreatedAt.AsTime(),
	}
	return res, nil
}
