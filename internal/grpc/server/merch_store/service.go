package merch_store

import (
	"Merch/internal/models"
	merch "Merch/pkg/api/v1"
	"context"
)

type (
	Authenticator interface {
		UserToken(ctx context.Context, credentials models.CredentialsDTO) (string, error)
		UserAuth(ctx context.Context, token string) (int64, error)
	}

	Repository interface {
		Info(ctx context.Context, userId int64) (models.UserInfo, error)
		BuyItem(ctx context.Context, userId int64, itemType string) error
		SendCoins(ctx context.Context, fromUserId int64, toUserLogin string, amount int64) error
	}
)

type Deps struct {
	Repository Repository
	Auth       Authenticator
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
