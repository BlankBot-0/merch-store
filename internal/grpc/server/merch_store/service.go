package merch_store

import (
	"Merch/internal/models"
	merch "Merch/pkg/api/v1"
	"context"
)

type (
	Shop interface {
		BuyItem(ctx context.Context, userId int64, itemType string) error
		SendCoins(ctx context.Context, fromUserId int64, toUserLogin string, amount int64) error
		Info(ctx context.Context, userId int64) (models.UserInfo, error)
	}
	Authenticator interface {
		UserToken(ctx context.Context, credentials models.CredentialsDTO) (string, error)
	}
)

type Deps struct {
	Shop Shop
	Auth Authenticator
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
