package service

import (
	"context"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	us "github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	u "github.com/gost1k337/url_shortener/user_service/api/protos/user"
	"time"
)

type UrlShort interface {
	Create(ctx context.Context, originalUrl string, expireAt time.Duration) (*CreateUrlShortResp, error)
	Get(ctx context.Context, token string) (*GetUrlShortResp, error)
}

type User interface {
	Create(ctx context.Context, username, email, password string) (*CreateUserResp, error)
	Get(ctx context.Context, id int64) (*GetUserResp, error)
	Delete(ctx context.Context, id int64) (*DeleteUserResp, error)
}

type Services struct {
	UrlShort
	User
}

type Deps struct {
	UrlShortService us.UrlShortsClient
	UserService     u.UserClient
}

func NewServices(deps *Deps, logger logging.Logger) *Services {
	return &Services{
		UrlShort: NewUrlShortService(deps.UrlShortService, logger),
		User:     NewUserService(deps.UserService, logger),
	}
}
