package urlshort_service

import (
	"fmt"

	g "github.com/gost1k337/url_shortener/api_gateway_service/pkg/grpc"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
)

func NewURLShortServiceConn(baseURL string, logger logging.Logger) (*g.BaseClient, error) {
	base, err := g.New(baseURL, logger)
	if err != nil {
		return nil, fmt.Errorf("conn: %w", err)
	}

	return base, nil
}
