package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gost1k337/url_shortener/api_gateway_service/config"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/client/urlshort_service"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/client/user_service"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/handlers"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/service"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/httpserver"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	us "github.com/gost1k337/url_shortener/url_shortening_service/api/protos/url_shorts"
	u "github.com/gost1k337/url_shortener/user_service/api/protos/user"
)

// Run
//
// @title           			Url Shortener
// @version         			1.0
// @description     			This is a REST API service for creating short urls.
// @host      					localhost:10000
// @BasePath  					/.
func Run(cfg *config.Config) {
	log := logging.NewLogger(cfg)

	log.Info("Connecting to user service...")

	uGrpcClient, err := user_service.NewUserServiceConn(
		fmt.Sprintf("%s:%s", cfg.UserService.Host, cfg.UserService.Port), log)
	if err != nil {
		log.Error(err)
	}

	uService := u.NewUserClient(uGrpcClient.GRPCClient)

	log.Info("Connecting to url-shortening service...")

	usGrpcClient, err := urlshort_service.NewURLShortServiceConn(
		fmt.Sprintf("%s:%s", cfg.URLShorteningService.Host, cfg.URLShorteningService.Port), log)
	if err != nil {
		log.Error(err)
	}

	usService := us.NewUrlShortsClient(usGrpcClient.GRPCClient)

	log.Info("Initializing services...")
	services := service.NewServices(&service.Deps{
		URLShortService: usService,
		UserService:     uService,
	}, log)

	log.Info("Initializing handlers...")
	handler := handlers.New(services, log, cfg)

	httpServer := httpserver.New(handler.HTTP(), httpserver.Port(cfg.App.Port))
	log.Infof("Server started on port %s", cfg.App.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	log.Info("Shutting down...")

	err = httpServer.Shutdown()

	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
