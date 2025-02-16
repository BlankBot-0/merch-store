package merch_store

import (
	merch "Merch/pkg/api/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Auth(ctx context.Context, request *merch.AuthRequest) (*merch.AuthResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.Deps.Auth.UserToken(ctx, request.GetPassword(), request.GetLogin())
	return &merch.AuthResponse{
		Token: token,
	}, err
}
