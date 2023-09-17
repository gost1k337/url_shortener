package grpc

import (
	"context"
	"errors"

	proto "github.com/gost1k337/url_shortener/user_service/api/protos/user"
	"github.com/gost1k337/url_shortener/user_service/internal/service"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserGrpcService struct {
	proto.UnimplementedUserServer
	logger logging.Logger
	u      service.User
}

func NewUserGrpcService(logger logging.Logger, su service.User) *UserGrpcService {
	return &UserGrpcService{
		proto.UnimplementedUserServer{},
		logger,
		su,
	}
}

func (s *UserGrpcService) Create(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	s.logger.Infof("User create method grpc call")

	user, err := s.u.Create(ctx, req.Username, req.Email, req.PasswordHash)
	if err != nil {
		s.logger.Error(err)

		return nil, status.Errorf(codes.Internal, "create user: %v", err)
	}

	resp := &proto.CreateUserResponse{
		Id:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    timestamppb.New(user.CreatedAt),
	}

	return resp, status.Errorf(codes.OK, "success")
}

func (s *UserGrpcService) Get(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	s.logger.Infof("User get method grpc call")

	user, err := s.u.GetByID(ctx, req.Id)
	if err != nil {
		s.logger.Error(err)

		if errors.Is(err, service.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "not found")
		}

		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	resp := &proto.GetUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}

	return resp, status.Errorf(codes.OK, "success")
}

func (s *UserGrpcService) Delete(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	s.logger.Infof("User delete method grpc call")

	user, err := s.u.Delete(ctx, req.Id)
	if err != nil {
		s.logger.Error(err)

		if errors.Is(err, service.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "delete user: %v", err)
	}

	resp := &proto.DeleteUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}

	return resp, status.Errorf(codes.OK, "success")
}
