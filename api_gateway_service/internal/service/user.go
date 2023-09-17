package service

import (
	"context"
	"fmt"

	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/hasher"
	"github.com/gost1k337/url_shortener/api_gateway_service/pkg/logging"
	u "github.com/gost1k337/url_shortener/user_service/api/protos/user"
	"google.golang.org/grpc/codes"
	status2 "google.golang.org/grpc/status"
)

type UserService struct {
	c      u.UserClient
	logger logging.Logger
}

func NewUserService(uc u.UserClient, logger logging.Logger) *UserService {
	return &UserService{
		c:      uc,
		logger: logger,
	}
}

func (s *UserService) Create(ctx context.Context, username, email, password string) (*CreateUserResp, error) {
	salt, err := hasher.GenerateRandomSalt(hasher.SaltSize)
	if err != nil {
		return nil, fmt.Errorf("generate salt: %w", err)
	}

	hash := hasher.HashPassword(password, salt)
	req := &u.CreateUserRequest{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
	}

	user, err := s.c.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("user create: %w", err)
	}

	resp := &CreateUserResp{
		ID:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.AsTime(),
	}

	return resp, nil
}

func (s *UserService) Get(ctx context.Context, id int64) (*GetUserResp, error) {
	req := &u.GetUserRequest{
		Id: id,
	}

	user, err := s.c.Get(ctx, req)
	if err != nil {
		status, ok := status2.FromError(err)
		if ok {
			if status.Code() == codes.NotFound {
				return nil, ErrUserNotFound
			}

			return nil, fmt.Errorf("grpc user get: %v, code (%s )", status.Message(), status.Code())
		}

		return nil, fmt.Errorf("non-grpc err: %w", err)
	}

	resp := &GetUserResp{
		ID:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.AsTime(),
	}

	return resp, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) (*DeleteUserResp, error) {
	req := &u.DeleteUserRequest{
		Id: id,
	}

	user, err := s.c.Delete(ctx, req)
	if err != nil {
		status, ok := status2.FromError(err)
		if ok {
			if status.Code() == codes.NotFound {
				return nil, ErrUserNotFound
			}

			return nil, fmt.Errorf("grpc user delete: %v, code (%s )", status.Message(), status.Code())
		}

		return nil, fmt.Errorf("non-grpc err: %w", err)
	}

	resp := &DeleteUserResp{
		ID:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.AsTime(),
	}

	return resp, nil
}
