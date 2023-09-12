package server

import (
	"fmt"
	proto "github.com/gost1k337/url_shortener/user_service/api/protos/user"
	"github.com/gost1k337/url_shortener/user_service/config"
	urlShortGrpc "github.com/gost1k337/url_shortener/user_service/internal/delivery/grpc"
	"github.com/gost1k337/url_shortener/user_service/internal/service"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

func NewUrlShortGrpcServer(cfg *config.Config, logger logging.Logger, services *service.Services) (func() error, *grpc.Server, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		return nil, nil, fmt.Errorf("net listen: %w", err)
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
	)

	urlShortGrpcService := urlShortGrpc.NewUserGrpcService(logger, services.User)
	proto.RegisterUserServer(grpcServer, urlShortGrpcService)

	if cfg.App.Debug {
		reflection.Register(grpcServer)
	}

	go func() {
		logger.Infof("UrlShort grpc server listening on port: %s", cfg.App.Port)
		logger.Fatal(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}