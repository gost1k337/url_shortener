package repository

import (
	"context"

	"github.com/gost1k337/url_shortener/user_service/internal/entity"
	"github.com/gost1k337/url_shortener/user_service/internal/repository/psql"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
	"github.com/gost1k337/url_shortener/user_service/pkg/postgres"
)

type User interface {
	Create(ctx context.Context, username, email, passwordHash string) (int64, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	Delete(ctx context.Context, id int64) error
}

type Repositories struct {
	User
}

func New(pg *postgres.Postgres, logger logging.Logger) *Repositories {
	return &Repositories{
		User: psql.NewUserRepo(pg, logger),
	}
}
