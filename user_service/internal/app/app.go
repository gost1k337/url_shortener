package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gost1k337/url_shortener/user_service/config"
	"github.com/gost1k337/url_shortener/user_service/internal/repository"
	"github.com/gost1k337/url_shortener/user_service/internal/service"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
	"github.com/gost1k337/url_shortener/user_service/pkg/postgres"
	"github.com/gost1k337/url_shortener/user_service/pkg/server"
)

func Run(cfg *config.Config) {
	log := logging.NewLogger(cfg)

	pg, err := postgres.New(cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Ping()
	if err != nil {
		log.Fatalf("db conn: %v", err)
	}

	log.Info("Postgres connected...")

	log.Info("Initializing repositories...")
	repos := repository.New(pg, log)

	log.Info("Initializing services...")
	services := service.NewServices(&service.ServicesDependencies{
		Repos: repos,
	}, log, cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	closeGrpcServer, grpcServer, err := server.NewUserGrpcServer(cfg, log, services)
	if err != nil {
		log.Error("new user grpc: %w", err)
	}

	defer func() {
		err = closeGrpcServer()
		if err != nil {
			log.Error("close grpc server: %w", err)
		}
	}()

	<-ctx.Done()
	grpcServer.GracefulStop()
}
