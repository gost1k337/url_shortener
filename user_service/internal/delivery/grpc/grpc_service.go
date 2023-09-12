package grpc

import (
	"context"
	proto "github.com/gost1k337/url_shortener/user_service/api/protos/user"
	"github.com/gost1k337/url_shortener/user_service/internal/service"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
)

type UserGrpcService struct {
	proto.UnimplementedUserServer
	logger logging.Logger
	su     service.User
}

func NewUserGrpcService(logger logging.Logger, su service.User) *UserGrpcService {
	return &UserGrpcService{
		logger: logger,
		su:     su,
	}
}

func (s *UserGrpcService) Create(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	//s.logger.Infof("UrlShort create method grpc call")
	//shortUrl, err := s.su.Create(ctx, 1, req.OriginalUrl, req.ExpireAt.AsTime())
	//if err != nil {
	//	s.logger.Error(err)
	//	return nil, fmt.Errorf("new short url: %w", err)
	//}
	//
	//resp := &proto.CreateUserResponse{
	//	Id:          shortUrl.Id,
	//	OriginalUrl: shortUrl.OriginalURL,
	//	ShortUrl:    shortUrl.ShortURL,
	//	ExpireAt:    timestamppb.New(shortUrl.ExpireAt),
	//	Visits:      shortUrl.Visits,
	//	CreatedAt:   timestamppb.New(shortUrl.CreatedAt),
	//}
	return nil, nil
}

func (s *UserGrpcService) Get(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	//s.logger.Infof("UrlShort get method grpc call")
	//shortUrl, err := s.su.GetByShortToken(ctx, req.Token)
	//if err != nil {
	//	s.logger.Error(err)
	//	return nil, fmt.Errorf("get short url: %w", err)
	//}
	//
	//resp := &proto.GetUserResponse{
	//	Id:          shortUrl.Id,
	//	OriginalUrl: shortUrl.OriginalURL,
	//	ShortUrl:    shortUrl.ShortURL,
	//	ExpireAt:    timestamppb.New(shortUrl.ExpireAt),
	//	Visits:      shortUrl.Visits,
	//	CreatedAt:   timestamppb.New(shortUrl.CreatedAt),
	//}
	//
	return nil, nil
}
