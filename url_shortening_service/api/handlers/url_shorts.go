package handlers

import (
	"context"
	"github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type UrlShorts struct {
	url_shorts.UnimplementedUrlShortsServer
	log logging.Logger
}

func NewUrlShorts(l logging.Logger) *UrlShorts {
	return &UrlShorts{log: l}
}

func (u *UrlShorts) CreateUrlShort(ctx context.Context, r *url_shorts.CreateUrlShortRequest) (*url_shorts.CreateUrlShortResponse, error) {
	u.log.Info("Handle CreateUrlShort:", zap.String("original url", r.GetOriginalUrl()), zap.Time("expire at", r.ExpireAt.AsTime()))

	return &url_shorts.CreateUrlShortResponse{
		OriginalUrl: "orig",
		ShortUrl:    "short",
		Visits:      0,
		ExpireAt:    timestamppb.New(time.Now()),
		CreatedAt:   timestamppb.New(time.Now()),
	}, nil
}
