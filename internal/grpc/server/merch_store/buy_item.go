package merch_store

import (
	"Merch/internal/auth"
	"Merch/internal/usecase/shop"
	merch "Merch/pkg/api/v1"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) BuyItem(ctx context.Context, request *merch.BuyItemRequest) (*emptypb.Empty, error) {
	userId := auth.GetUserIDFromCtx(ctx)
	err := s.Deps.Shop.BuyItem(ctx, userId, request.GetItem())
	if errors.Is(err, shop.ErrItemIsNotFound) {
		err = status.Errorf(codes.InvalidArgument, "item is not found")
	} else if errors.Is(err, shop.ErrNotEnoughCoins) {
		err = status.Errorf(codes.FailedPrecondition, "not enough coins")
	} else if err != nil {
		err = status.Errorf(codes.Internal, "internal server error")
	}
	return &emptypb.Empty{}, err
}
