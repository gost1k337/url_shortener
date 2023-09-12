package service

import (
	"context"
	"github.com/gost1k337/url_shortener/user_service/config"
	"github.com/gost1k337/url_shortener/user_service/internal/entity"
	"github.com/gost1k337/url_shortener/user_service/internal/repository"
	"github.com/gost1k337/url_shortener/user_service/internal/repository/repoerrors"
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

func (s *UserService) Create(ctx context.Context, username, email, passwordHash string) (*entity.User, error) {
	id, err := s.repo.Create(ctx, username, email, passwordHash)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetById(ctx context.Context, id int64) (*entity.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == repoerrors.ErrNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) (*entity.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		if err == repoerrors.ErrNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
