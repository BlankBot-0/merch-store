package merch_store

import (
	"Merch/internal/auth"
	merch "Merch/pkg/api/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) SendCoin(ctx context.Context, request *merch.SendCoinRequest) (*emptypb.Empty, error) {
	userId := auth.GetUserIDFromCtx(ctx)
	err := s.Deps.Shop.SendCoins(ctx, userId, request.GetToUser(), request.GetAmount())
	return &emptypb.Empty{}, err
}
