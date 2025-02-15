package merch_platform

import (
	"Merch/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

type (
	// ROSalesPlatform is a read-only repository
	ROSalesPlatform interface {
		UserById(ctx context.Context, userId int64) (models.User, error)
		UserByLogin(ctx context.Context, login string) (models.User, error)
		UserInventory(ctx context.Context, userId int64) ([]models.InventoryItem, error)
		Item(ctx context.Context, itemType string) (models.Item, error)
		UserSendOperations(ctx context.Context, userId int64) ([]models.Transaction, error)
		UserReceivedOperations(ctx context.Context, userId int64) ([]models.Transaction, error)
	}

	// RWSalesPlatform is a read-write repository
	RWSalesPlatform interface {
		RemoveCoins(ctx context.Context, userId, amount int64) error
		AddCoins(ctx context.Context, userId, amount int64) error
		AddTransaction(ctx context.Context, fromUserId, toUserId int64, amount int64) error
		AddPurchase(ctx context.Context, userId, itemId int64) error
		ROSalesPlatform
	}

	RepositoryProvider interface {
		ROSalesPlatform() ROSalesPlatform
		RWSalesPlatform() RWSalesPlatform
	}

	DB interface {
		RepositoryProvider
		RunInTx(ctx context.Context, f func(tx RepositoryProvider) error, isoLevel pgx.TxIsoLevel) error
	}
)

type Deps struct {
	repo DB
}
type MerchPlatform struct {
	Deps Deps
}

func New(deps Deps) *MerchPlatform {
	return &MerchPlatform{
		Deps: deps,
	}
}

// BuyItem
func (s *MerchPlatform) BuyItem(ctx context.Context, userId int64, itemType string) error {
	err := s.Deps.repo.RunInTx(ctx, func(tx RepositoryProvider) error {
		user, err := tx.ROSalesPlatform().UserById(ctx, userId)
		if err != nil {
			return err
		}
		item, err := tx.ROSalesPlatform().Item(ctx, itemType)
		if err != nil {
			return err
		}

		if user.Coins < item.Coins {
			return ErrNotEnoughCoins
		}

		err = tx.RWSalesPlatform().RemoveCoins(ctx, userId, item.Coins)
		if err != nil {
			return err
		}

		return tx.RWSalesPlatform().AddPurchase(ctx, userId, item.Id)
	}, pgx.Serializable)

	return err
}

func (s *MerchPlatform) SendCoins(ctx context.Context, fromUserId int64, toUserLogin string, amount int64) error {
	if amount <= 0 {
		return ErrIncorrectAmount
	}

	err := s.Deps.repo.RunInTx(ctx, func(tx RepositoryProvider) error {
		sender, err := tx.ROSalesPlatform().UserById(ctx, fromUserId)
		if err != nil {
			return err
		}

		if sender.Coins < amount {
			return ErrNotEnoughCoins
		}

		receiver, err := tx.ROSalesPlatform().UserByLogin(ctx, toUserLogin)
		if err != nil {
			return err
		}

		err = tx.RWSalesPlatform().RemoveCoins(ctx, fromUserId, amount)
		if err != nil {
			return err
		}

		return tx.RWSalesPlatform().AddCoins(ctx, receiver.Id, amount)
	}, pgx.Serializable)

	return err
}

func (s *MerchPlatform) Info(ctx context.Context, userId int64) (models.UserInfo, error) {
	var userInfo models.UserInfo

	err := s.Deps.repo.RunInTx(ctx, func(tx RepositoryProvider) error {
		user, err := tx.ROSalesPlatform().UserById(ctx, userId)
		if err != nil {
			return err
		}

		coinsSentHistory, err := tx.ROSalesPlatform().UserSendOperations(ctx, userId)
		if err != nil {
			return err
		}
		coinsReceivedHistory, err := tx.ROSalesPlatform().UserReceivedOperations(ctx, userId)
		if err != nil {
			return err
		}

		inventory, err := tx.ROSalesPlatform().UserInventory(ctx, userId)

		userInfo = models.UserInfo{
			Coins:     user.Coins,
			Inventory: inventory,
			Sent:      coinsSentHistory,
			Received:  coinsReceivedHistory,
		}

		return err
	}, pgx.ReadCommitted)

	return userInfo, err
}
