package merch_store

import (
	"Merch/internal/auth"
	merch "Merch/pkg/api/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) BuyItem(ctx context.Context, request *merch.BuyItemRequest) (*emptypb.Empty, error) {
	userId := auth.GetUserIDFromCtx(ctx)
	err := s.Deps.Shop.BuyItem(ctx, userId, request.GetItem())
	return &emptypb.Empty{}, err
}
