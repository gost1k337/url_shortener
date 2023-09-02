package app

import (
	"fmt"
	"github.com/gost1k337/url_shortener/api_gateway_service/config"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/handlers"
	"github.com/gost1k337/url_shortener/api_gateway_service/internal/service"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/httpserver"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	log := logging.NewLogger(cfg)

	log.Info("Initializing services...")
	services := service.NewServices(&service.ServicesDependencies{}, log)

	log.Info("Initializing handlers...")
	handler := handlers.New(services, log)

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
	err := httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
