package service

import (
	"context"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	us "github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	"time"
)

type UrlShort interface {
	Create(ctx context.Context, originalUrl string, expireAt time.Duration) (*CreateUrlShortResp, error)
	Get(ctx context.Context, token string) (*GetUrlShortResp, error)
}

type Services struct {
	UrlShort
}

type Deps struct {
	UrlShortService us.UrlShortsClient
}

func NewServices(deps *Deps, logger logging.Logger) *Services {
	return &Services{
		UrlShort: NewUrlShortService(deps.UrlShortService, logger),
	}
}
