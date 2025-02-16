package merch_store

import (
	"Merch/internal/models"
	merch "Merch/pkg/api/v1"
	"context"
)

type (
	shopService interface {
		BuyItem(ctx context.Context, userId int64, itemType string) error
		SendCoins(ctx context.Context, fromUserId int64, toUserLogin string, amount int64) error
		Info(ctx context.Context, userId int64) (*models.UserInfo, error)
	}
	authService interface {
		UserToken(ctx context.Context, login, password string) (string, error)
	}
)

type Deps struct {
	Shop shopService
	Auth authService
}

type Service struct {
	merch.MerchStoreServer
	Deps Deps
}

func NewService(deps Deps) *Service {
	return &Service{
		Deps: deps,
	}
}
