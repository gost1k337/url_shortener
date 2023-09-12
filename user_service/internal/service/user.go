package service

import (
	"context"
	"github.com/gost1k337/url_shortener/user_service/config"
	"github.com/gost1k337/url_shortener/user_service/internal/entity"
	"github.com/gost1k337/url_shortener/user_service/internal/repository"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
)

type UserService struct {
	repo   repository.User
	logger logging.Logger
	cfg    *config.Config
}

func NewUserService(repo repository.User, logger logging.Logger, cfg *config.Config) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
		cfg:    cfg,
	}
}

func (s *UserService) Create(ctx context.Context, username, string, passwordHash string) (*entity.User, error) {
	return nil, nil
}

func (s *UserService) GetById(ctx context.Context, id int64) (*entity.User, error) {
	return nil, nil
}
