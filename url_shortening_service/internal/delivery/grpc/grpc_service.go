package grpc

import (
	"context"
	"fmt"

	proto "github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/service"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UrlShortGrpcService struct {
	proto.UnimplementedUrlShortsServer
	logger logging.Logger
	su     service.ShortURL
}

func NewUrlShortGrpcService(logger logging.Logger, su service.ShortURL) *UrlShortGrpcService {
	return &UrlShortGrpcService{
		logger: logger,
		su:     su,
	}
}

func (s *UrlShortGrpcService) Create(ctx context.Context, req *proto.CreateUrlShortRequest) (
	*proto.CreateUrlShortResponse, error,
) {
	s.logger.Infof("UrlShort create method grpc call")

	shortUrl, err := s.su.Create(ctx, 1, req.OriginalUrl, req.ExpireAt.AsTime())
	if err != nil {
		s.logger.Error(err)

		return nil, fmt.Errorf("new short url: %w", err)
	}

	resp := &proto.CreateUrlShortResponse{
		Id:          shortUrl.Id,
		OriginalUrl: shortUrl.OriginalURL,
		ShortUrl:    shortUrl.ShortURL,
		ExpireAt:    timestamppb.New(shortUrl.ExpireAt),
		Visits:      shortUrl.Visits,
		CreatedAt:   timestamppb.New(shortUrl.CreatedAt),
	}

	return resp, nil
}

func (s *UrlShortGrpcService) Get(ctx context.Context, req *proto.GetUrlShortRequest) (
	*proto.GetUrlShortResponse, error,
) {
	s.logger.Infof("UrlShort get method grpc call")

	shortUrl, err := s.su.GetByShortToken(ctx, req.Token)
	if err != nil {
		s.logger.Error(err)

		return nil, fmt.Errorf("get short url: %w", err)
	}

	resp := &proto.GetUrlShortResponse{
		Id:          shortUrl.Id,
		OriginalUrl: shortUrl.OriginalURL,
		ShortUrl:    shortUrl.ShortURL,
		ExpireAt:    timestamppb.New(shortUrl.ExpireAt),
		Visits:      shortUrl.Visits,
		CreatedAt:   timestamppb.New(shortUrl.CreatedAt),
	}

	return resp, nil
}
