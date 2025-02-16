package postgres

import (
	"Merch/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

var _ DB = (*Database)(nil)

type (
	// ROShop is a read-only repository
	ROShop interface {
		UserInventory(ctx context.Context, userId int64) ([]models.InventoryItem, error)
		GetItem(ctx context.Context, itemType string) (models.Item, error)
		UserTransactions(ctx context.Context, userId int64) ([]models.Transaction, error)
	}

	// RWShop is a read-write repository
	RWShop interface {
		AddCoins(ctx context.Context, userId, amount int64) error
		SaveTransaction(ctx context.Context, fromUserId, toUserId int64, amount int64) error
		AddPurchase(ctx context.Context, userId, itemId int64) error
		ROShop
	}

	ROUsers interface {
		UserById(ctx context.Context, userId int64) (models.User, error)
		UserByLogin(ctx context.Context, login string) (models.User, error)
	}

	RWUsers interface {
		CreateUser(ctx context.Context, login, passwordHash string) (int64, error)
		ROUsers
	}

	RepositoryProvider interface {
		ROShop() ROShop
		RWShop() RWShop
		ROUsers() ROUsers
		RWUsers() RWUsers
	}

	DB interface {
		RepositoryProvider
		RunInTx(ctx context.Context, f func(tx RepositoryProvider) error, isoLevel pgx.TxIsoLevel) error
	}
)
