package shop

import (
	"Merch/internal/models"
	"Merch/internal/postgres"
	"context"
	"errors"
	"fmt"
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
	item, err := s.Deps.Repo.ROShop().GetItem(ctx, itemType)
	if errors.Is(err, postgres.ErrNotFound) {
		return ErrItemIsNotFound
	} else if err != nil {
		return fmt.Errorf("unexpected error occured while getting user: %w", err)
	}

	return s.Deps.Repo.RunInTx(ctx, func(tx postgres.RepositoryProvider) error {
		user, err := tx.ROUsers().UserById(ctx, userId)
		if errors.Is(err, postgres.ErrNotFound) {
			return ErrUserIsNotFound
		} else if err != nil {
			return fmt.Errorf("unexpected error occured while getting user: %w", err)
		}

		if user.Coins < item.Coins {
			return ErrNotEnoughCoins
		}

		if err = tx.RWShop().AddCoins(ctx, userId, -item.Coins); errors.Is(err, postgres.ErrNotChanged) {
			return fmt.Errorf("failed to withdrawing coins from user %d: %w", user.Id, err)
		} else if err != nil {
			return fmt.Errorf("error occured while withdrawing coins from user %d: %w", user.Id, err)
		}

		if err := tx.RWShop().AddPurchase(ctx, userId, item.Id); errors.Is(err, postgres.ErrNotChanged) {
			return fmt.Errorf("failed to save purchase information for user %d: %w", user.Id, err)
		} else if err != nil {
			return fmt.Errorf("error occured while saving purchase info for user %d: %w", user.Id, err)
		}

		return nil
	}, pgx.Serializable)
}

func (s *Shop) SendCoins(ctx context.Context, fromUserId int64, toUserLogin string, amount int64) error {
	receiver, err := s.Deps.Repo.ROUsers().UserByLogin(ctx, toUserLogin)
	if errors.Is(err, postgres.ErrNotFound) {
		return ErrUserIsNotFound
	} else if err != nil {
		return fmt.Errorf("unexpected error occured while getting receiver user by login: %w", err)
	}

	return s.Deps.Repo.RunInTx(ctx, func(tx postgres.RepositoryProvider) error {
		sender, err := tx.ROUsers().UserById(ctx, fromUserId)
		if errors.Is(err, postgres.ErrNotFound) {
			return ErrUserIsNotFound
		} else if err != nil {
			return fmt.Errorf("unexpected error occured while getting sender user by id: %w", err)
		}

		if sender.Coins < amount {
			return ErrNotEnoughCoins
		}

		if err = tx.RWShop().AddCoins(ctx, fromUserId, -amount); err != nil {
			return fmt.Errorf("unexpected error occured while withdrawing coins: %w", err)
		}

		if err = tx.RWShop().AddCoins(ctx, receiver.Id, amount); err != nil {
			return fmt.Errorf("unexpected error occured while adding coins: %w", err)
		}

		if err = tx.RWShop().SaveTransaction(ctx, fromUserId, receiver.Id, amount); err != nil {
			return fmt.Errorf("unexpected error occured while saving transaction details: %w", err)
		}

		return nil
	}, pgx.Serializable)
}

func (s *Shop) Info(ctx context.Context, userId int64) (*models.UserInfo, error) {

	roUsers := s.Deps.Repo.ROUsers()
	roShop := s.Deps.Repo.ROShop()

	user, err := roUsers.UserById(ctx, userId)
	if errors.Is(err, postgres.ErrNotFound) {
		return nil, ErrUserIsNotFound
	} else if err != nil {
		return nil, fmt.Errorf("failed to resolve user by ID: %w", err)
	}

	coinsHistory, err := roShop.UserTransactions(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to list user trasactions: %w", err)
	}

	inventory, err := roShop.UserInventory(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to list user inventory: %w", err)
	}

	return &models.UserInfo{
		Coins:     user.Coins,
		Inventory: inventory,
		Sent: lo.Filter(coinsHistory, func(item models.Transaction, _ int) bool {
			return item.Sender == user.Login
		}),
		Received: lo.Filter(coinsHistory, func(item models.Transaction, _ int) bool {
			return item.Receiver == user.Login
		}),
	}, nil
}
