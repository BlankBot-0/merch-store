package merch_store

import (
	"Merch/internal/models"
	merch "Merch/pkg/api/v1"
	"context"
)

func (s *Service) Auth(ctx context.Context, request *merch.AuthRequest) (*merch.AuthResponse, error) {
	token, err := s.Deps.Auth.UserToken(ctx, models.CredentialsDTO{
		Password: request.GetPassword(),
		Username: request.GetLogin(),
	})
	return &merch.AuthResponse{
		Token: token,
	}, err
}
