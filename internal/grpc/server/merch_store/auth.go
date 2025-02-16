package merch_store

import (
	"Merch/internal/usecase/auth"
	merch "Merch/pkg/api/v1"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Auth(ctx context.Context, request *merch.AuthRequest) (*merch.AuthResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.Deps.Auth.UserToken(ctx, request.GetLogin(), request.GetPassword())
	if errors.Is(err, auth.ErrUserAlreadyExists) {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &merch.AuthResponse{
		Token: token,
	}, err
}
