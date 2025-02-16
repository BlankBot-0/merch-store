package shop

import (
	"Merch/internal/models"
	"Merch/internal/postgres"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
)

type Deps struct {
	Repo postgres.DB
}
type Shop struct {
	Deps Deps
}

func New(deps Deps) *Shop {
	return &Shop{
		Deps: deps,
	}
}

func (s *Shop) BuyItem(ctx context.Context, userId int64, itemType string) error {
	err := s.Deps.Repo.RunInTx(ctx, func(tx postgres.RepositoryProvider) error {
		user, err := tx.ROUsers().UserById(ctx, userId)
		if err != nil {
			return err
		}
		item, err := tx.ROShop().GetItem(ctx, itemType)
		if err != nil {
			return err
		}

		if user.Coins < item.Coins {
			return ErrNotEnoughCoins
		}

		err = tx.RWShop().AddCoins(ctx, userId, -item.Coins)
		if err != nil {
			return err
		}

		return tx.RWShop().AddPurchase(ctx, userId, item.Id)
	}, pgx.Serializable)

	return err
}

func (s *Shop) SendCoins(ctx context.Context, fromUserId int64, toUserLogin string, amount int64) error {
	if amount <= 0 {
		return ErrIncorrectAmount
	}

	err := s.Deps.Repo.RunInTx(ctx, func(tx postgres.RepositoryProvider) error {
		sender, err := tx.ROUsers().UserById(ctx, fromUserId)
		if err != nil {
			return err
		}

		if sender.Coins < amount {
			return ErrNotEnoughCoins
		}

		receiver, err := tx.ROUsers().UserByLogin(ctx, toUserLogin)
		if err != nil {
			return err
		}

		err = tx.RWShop().AddCoins(ctx, fromUserId, -amount)
		if err != nil {
			return err
		}

		err = tx.RWShop().AddCoins(ctx, receiver.Id, amount)
		if err != nil {
			return err
		}

		err = tx.RWShop().SaveTransaction(ctx, fromUserId, receiver.Id, amount)
		return err
	}, pgx.Serializable)

	return err
}

func (s *Shop) Info(ctx context.Context, userId int64) (models.UserInfo, error) {
	var userInfo models.UserInfo

	err := s.Deps.Repo.RunInTx(ctx, func(tx postgres.RepositoryProvider) error {
		user, err := tx.ROUsers().UserById(ctx, userId)
		if err != nil {
			return err
		}

		coinsHistory, err := tx.ROShop().UserTransactions(ctx, userId)
		if err != nil {
			return err
		}

		inventory, err := tx.ROShop().UserInventory(ctx, userId)
		if err != nil {
			return err
		}

		userInfo = models.UserInfo{
			Coins:     user.Coins,
			Inventory: inventory,
			Sent: lo.Filter(coinsHistory, func(item models.Transaction, _ int) bool {
				return item.Sender == user.Login
			}),
			Received: lo.Filter(coinsHistory, func(item models.Transaction, _ int) bool {
				return item.Receiver == user.Login
			}),
		}

		return nil
	}, pgx.ReadCommitted)

	return userInfo, err
}
