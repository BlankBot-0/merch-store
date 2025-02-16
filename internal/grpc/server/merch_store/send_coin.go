package merch_store

import (
	"Merch/internal/auth"
	merch "Merch/pkg/api/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) SendCoin(ctx context.Context, request *merch.SendCoinRequest) (*emptypb.Empty, error) {
	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userId := auth.GetUserIDFromCtx(ctx)
	err := s.Deps.Shop.SendCoins(ctx, userId, request.GetToUser(), request.GetAmount())
	return &emptypb.Empty{}, err
}
