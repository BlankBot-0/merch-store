package merch_store

import (
	"Merch/internal/grpc/server/merch_store/fake"
	"Merch/internal/models"
	pb "Merch/pkg/api/v1"
	"context"
	"testing"
)

var (
	fakeInfo = models.UserInfo{
		Coins: 500,
		Inventory: []models.InventoryItem{
			{
				Type:     "type1",
				Quantity: 5,
			},
		},
		Sent: []models.Transaction{
			{
				Receiver: "vanya",
				Sender:   "petya",
				Amount:   300,
			},
		},
		Received: []models.Transaction{
			{
				Receiver: "petya",
				Sender:   "masha",
				Amount:   200,
			},
		},
	}
)

func TestService_Shop_Info(t *testing.T) {
	fakeShop := fake.ShopServiceFake{
		FakeInfo: fakeInfo,
	}

	service := NewService(Deps{
		Shop: &fakeShop,
	})

	resp, err := service.Info(context.Background(), &pb.InfoRequest{})
	if err != nil {
		t.Fatalf("got unexpected error: %s", err)
	}

	if resp.Coins != 500 {
		t.Errorf("got %d coins, want 500", resp.Coins)
	}
	if len(resp.Inventory) != 1 {
		t.Errorf("got %d items, want 1", len(resp.Inventory))
	}
	if len(resp.CoinHistory.Sent) != 1 {
		t.Errorf("got %d sending transactions, want 1", len(resp.CoinHistory.Sent))
	}
	if len(resp.CoinHistory.Received) != 1 {
		t.Errorf("got %d receiving transactions, want 1", len(resp.CoinHistory.Received))
	}
}
