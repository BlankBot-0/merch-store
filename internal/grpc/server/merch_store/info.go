package merch_store

import (
	"Merch/internal/auth"
	"Merch/internal/models"
	merch "Merch/pkg/api/v1"
	"context"
	"github.com/samber/lo"
)

func (s *Service) Info(ctx context.Context, request *merch.InfoRequest) (*merch.InfoResponse, error) {
	res, err := s.Deps.Shop.Info(ctx, auth.GetUserIDFromCtx(ctx))
	if err != nil {
		return nil, err
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
