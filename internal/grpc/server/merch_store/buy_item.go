package merch_store

import (
	merch "Merch/pkg/api/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) BuyItem(ctx context.Context, request *merch.BuyItemRequest) (*emptypb.Empty, error) {
	err := s.Deps.Repository.BuyItem(ctx, 0, request.GetItem())
	return &emptypb.Empty{}, err
}
