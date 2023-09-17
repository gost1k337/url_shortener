package service

import (
	"context"
	"errors"
	"fmt"

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
		return nil, fmt.Errorf("create user: %w", err)
	}

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, fmt.Errorf("get user err: %w", err)
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) (*entity.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, fmt.Errorf("get user err: %w", err)
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("delete user err: %w", err)
	}

	return user, nil
}
