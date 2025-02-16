package shop

import (
	"Merch/internal/models"
	"Merch/internal/postgres"
	"Merch/internal/postgres/fake"
	"context"
	"errors"
	"testing"
)

var someErr = errors.New("some error")

func TestShop_BuyItem(t *testing.T) {
	ctx := context.Background()

	items := map[string]models.Item{
		"item": {
			Id:    1,
			Type:  "item",
			Coins: 20,
		},
	}

	testCases := []struct {
		name        string
		db          fake.DBMock
		expectedErr error
	}{
		{
			name: "ok",
			db: fake.DBMock{
				UsersRepoFake: &fake.UsersRepoFake{},
				ShopRepoFake: &fake.ShopRepoFake{
					Items: items,
				},
			},
		},
		{
			name: "user not found",
			db: fake.DBMock{
				UsersRepoFake: &fake.UsersRepoFake{
					UserByIdErr: postgres.ErrNotFound,
				},
				ShopRepoFake: &fake.ShopRepoFake{
					Items: items,
				},
			},
			expectedErr: ErrUserIsNotFound,
		},
		{
			name: "can't add coins",
			db: fake.DBMock{
				UsersRepoFake: &fake.UsersRepoFake{},
				ShopRepoFake: &fake.ShopRepoFake{
					Items:       items,
					AddCoinsErr: postgres.ErrNotChanged,
				},
			},
			expectedErr: postgres.ErrNotChanged,
		},
		{
			name: "can't add purchase",
			db: fake.DBMock{
				UsersRepoFake: &fake.UsersRepoFake{},
				ShopRepoFake: &fake.ShopRepoFake{
					Items:          items,
					AddPurchaseErr: postgres.ErrNotChanged,
				},
			},
			expectedErr: postgres.ErrNotChanged,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, _ := tc.db.CreateUser(ctx, "login", "password_hash")

			shop := New(Deps{Repo: tc.db})

			err := shop.BuyItem(ctx, id, "item")
			if tc.expectedErr == nil && err != nil {
				t.Fatalf("got unexpected error")
			} else if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("expected %s error, got %s", tc.expectedErr, err)
			}
		})
	}
}

func TestShop_SendCoins(t *testing.T) {
	ctx := context.Background()

	items := map[string]models.Item{
		"item": {
			Id:    1,
			Type:  "item",
			Coins: 20,
		},
	}

	testCases := []struct {
		name        string
		db          fake.DBMock
		expectedErr error
	}{
		{
			name: "ok",
			db: fake.DBMock{
				UsersRepoFake: &fake.UsersRepoFake{},
				ShopRepoFake: &fake.ShopRepoFake{
					Items: items,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, _ := tc.db.CreateUser(ctx, "login1", "password_hash")
			_, _ = tc.db.CreateUser(ctx, "login2", "password_hash")

			shop := New(Deps{Repo: tc.db})

			err := shop.SendCoins(ctx, id, "login2", 20)
			if tc.expectedErr == nil && err != nil {
				t.Fatalf("got unexpected error")
			} else if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("expected %s error, got %s", tc.expectedErr, err)
			}
		})
	}
}

func TestShop_Info(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name        string
		db          fake.DBMock
		expectedErr error
	}{
		{
			name: "ok",
			db: fake.DBMock{
				UsersRepoFake: &fake.UsersRepoFake{},
				ShopRepoFake: &fake.ShopRepoFake{
					Transactions: map[int64][]models.Transaction{
						1: {
							models.Transaction{
								Receiver: "login1",
								Sender:   "login2",
								Amount:   20,
							},
							models.Transaction{
								Receiver: "login2",
								Sender:   "login1",
								Amount:   15,
							},
						},
					},
					Inventories: map[int64][]models.InventoryItem{
						1: {
							models.InventoryItem{
								Type:     "item",
								Quantity: 5,
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, _ := tc.db.CreateUser(ctx, "login1", "password_hash")
			_, _ = tc.db.CreateUser(ctx, "login1", "password_hash")

			shop := New(Deps{Repo: tc.db})

			info, err := shop.Info(ctx, id)
			if tc.expectedErr == nil && err != nil {
				t.Fatalf("got unexpected error")
			} else if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("expected %s error, got %s", tc.expectedErr, err)
			}

			// todo validate info
			_ = info
		})
	}
}
