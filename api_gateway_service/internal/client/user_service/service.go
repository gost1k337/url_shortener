package user_service

import (
	g "github.com/gost1k337/url_shortener/api_gateway_service/pkg/grpc"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
)

func NewUserServiceConn(baseUrl string, logger logging.Logger) (*g.BaseClient, error) {
	base, err := g.New(baseUrl, logger)
	if err != nil {
		return nil, err
	}
	return base, nil
}
