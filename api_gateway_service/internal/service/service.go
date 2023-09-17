package service

import (
	"context"
	"time"

	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	us "github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	u "github.com/gost1k337/url_shortener/user_service/api/protos/user"
)

type URLShort interface {
	Create(ctx context.Context, originalURL string, expireAt time.Duration) (*CreateURLShortResp, error)
	Get(ctx context.Context, token string) (*GetURLShortResp, error)
}

type User interface {
	Create(ctx context.Context, username, email, password string) (*CreateUserResp, error)
	Get(ctx context.Context, id int64) (*GetUserResp, error)
	Delete(ctx context.Context, id int64) (*DeleteUserResp, error)
}

type Services struct {
	URLShort
	User
}

type Deps struct {
	URLShortService us.UrlShortsClient
	UserService     u.UserClient
}

func NewServices(deps *Deps, logger logging.Logger) *Services {
	return &Services{
		URLShort: NewURLShortService(deps.URLShortService, logger),
		User:     NewUserService(deps.UserService, logger),
	}
}
