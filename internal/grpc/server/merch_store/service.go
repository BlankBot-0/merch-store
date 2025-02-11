package merch_store

import "context"

type Repository interface {
	Info(ctx context.Context, userId int) (indoDTO, error)
	SendCoin(ctx context.Context, fromUserId, toUserId, coins int) error
	BuyItem(ctx context.Context, userId, itemId int) error
}

type Auth interface {
	Token(ctx context.Context, credentials CredentialsDTO) (string, error)
	Auth(ctx context.Context, token string) error
}

type Deps struct {
	Repository Repository
	Auth       Auth
}

type Service struct {
	app.Server
	Deps Deps
}

func NewService(deps Deps) *Service {
	return &Service{
		Deps: deps,
	}
}
