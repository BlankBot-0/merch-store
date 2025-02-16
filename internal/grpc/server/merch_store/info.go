package merch_store

import (
	"Merch/internal/auth"
	"Merch/internal/models"
	"Merch/internal/usecase/shop"
	merch "Merch/pkg/api/v1"
	"context"
	"errors"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Info(ctx context.Context, request *merch.InfoRequest) (*merch.InfoResponse, error) {
	res, err := s.Deps.Shop.Info(ctx, auth.GetUserIDFromCtx(ctx))
	if errors.Is(err, shop.ErrUserIsNotFound) {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &merch.InfoResponse{
		Coins: res.Coins,
		Inventory: lo.Map(res.Inventory, func(item models.InventoryItem, index int) *merch.InfoResponseItem {
			return &merch.InfoResponseItem{
				Type:     item.Type,
				Quantity: item.Quantity,
			}
		}),
		CoinHistory: &merch.InfoResponseCoinHistoryMessage{
			Sent: lo.Map(res.Sent, func(t models.Transaction, index int) *merch.InfoResponseCoinHistoryMessageSendCoinEntry {
				return &merch.InfoResponseCoinHistoryMessageSendCoinEntry{
					ToUser: t.Receiver,
					Amount: t.Amount,
				}
			}),
			Received: lo.Map(res.Received, func(t models.Transaction, index int) *merch.InfoResponseCoinHistoryMessageReceiveCoinEntry {
				return &merch.InfoResponseCoinHistoryMessageReceiveCoinEntry{
					FromUser: t.Sender,
					Amount:   t.Amount,
				}
			}),
		},
	}, err
}
