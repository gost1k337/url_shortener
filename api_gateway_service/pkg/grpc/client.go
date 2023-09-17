package grpc

import (
	"fmt"

	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BaseClient struct {
	BaseURL    string
	GRPCClient *grpc.ClientConn
	logger     logging.Logger
}

func New(baseURL string, logger logging.Logger) (*BaseClient, error) {
	conn, err := grpc.Dial(baseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc: %w", err)
	}

	client := &BaseClient{
		BaseURL:    baseURL,
		GRPCClient: conn,
		logger:     logger,
	}

	return client, nil
}
