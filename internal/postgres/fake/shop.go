package fake

import (
	"Merch/internal/models"
	"Merch/internal/postgres"
	"context"
)

var _ postgres.RWShop = &ShopRepoFake{}

type ShopRepoFake struct {
	Items        map[string]models.Item
	Transactions map[int64][]models.Transaction
	Inventories  map[int64][]models.InventoryItem
	Coins        map[int64]int64

	AddCoinsErr         error
	SaveTransactionErr  error
	AddPurchaseErr      error
	UserInventoryErr    error
	GetItemErr          error
	UserTransactionsErr error
}

func (s ShopRepoFake) AddCoins(_ context.Context, _, _ int64) error {
	return s.AddCoinsErr
}

func (s ShopRepoFake) SaveTransaction(_ context.Context, _, _ int64, _ int64) error {
	return s.SaveTransactionErr
}

func (s ShopRepoFake) AddPurchase(_ context.Context, _, _ int64) error {
	return s.AddPurchaseErr
}

func (s ShopRepoFake) UserInventory(_ context.Context, userId int64) ([]models.InventoryItem, error) {
	if s.UserInventoryErr != nil {
		return nil, s.UserInventoryErr
	}
	return s.Inventories[userId], nil
}

func (s ShopRepoFake) GetItem(_ context.Context, itemType string) (models.Item, error) {
	if s.GetItemErr != nil {
		return models.Item{}, s.GetItemErr
	}

	return s.Items[itemType], nil
}

func (s ShopRepoFake) UserTransactions(_ context.Context, userId int64) ([]models.Transaction, error) {
	if s.UserTransactionsErr != nil {
		return nil, s.UserTransactionsErr
	}

	return s.Transactions[userId], nil
}
