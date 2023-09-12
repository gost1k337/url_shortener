package service

import (
	"context"
	"github.com/gost1k337/url_shortener/user_service/config"
	"github.com/gost1k337/url_shortener/user_service/internal/entity"
	"github.com/gost1k337/url_shortener/user_service/internal/repository"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
)

type User interface {
	Create(ctx context.Context, username, email, passwordHash string) (*entity.User, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) (*entity.User, error)
}

type Services struct {
	User User
}

type ServicesDependencies struct {
	Repos *repository.Repositories
}

func NewServices(deps *ServicesDependencies, logger logging.Logger, cfg *config.Config) *Services {
	return &Services{
		User: NewUserService(deps.Repos.User, logger, cfg),
	}
}
