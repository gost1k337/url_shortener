package app

import (
	"fmt"
	"github.com/gost1k337/url_shortener/url_shortening_service/config"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/handlers"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/repository"
	"github.com/gost1k337/url_shortener/url_shortening_service/internal/service"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/httpserver"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/logging"
	"github.com/gost1k337/url_shortener/url_shortening_service/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	log := logging.NewLogger(cfg)

	pg, err := postgres.New(cfg.Db.DSN)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Postgres connected...")

	log.Info("Initializing repositories...")
	repos := repository.New(pg, log)

	log.Info("Initializing services...")
	services := service.NewServices(&service.ServicesDependencies{
		Repos: repos,
	}, log, cfg)

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
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
