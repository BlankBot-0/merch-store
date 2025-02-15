package merch_store

import (
	merch "Merch/pkg/api/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) SendCoin(ctx context.Context, request *merch.SendCoinRequest) (*emptypb.Empty, error) {
	err := s.Deps.Repository.SendCoins(ctx, 0, request.GetToUser(), request.GetAmount())
	return &emptypb.Empty{}, err
}
