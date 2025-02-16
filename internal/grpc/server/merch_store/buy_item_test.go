package merch_store

import (
	"Merch/internal/auth"
	"Merch/internal/grpc/server/merch_store/fake"
	"Merch/internal/models"
	pb "Merch/pkg/api/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var (
	fakeItemsByType = map[string]models.Item{
		"bipki": {
			Id:    1,
			Type:  "bipki",
			Coins: 1000,
		},
		"marshmallow": {
			Id:    2,
			Type:  "marshmallow",
			Coins: 10,
		},
	}
)

func TestService_BuyItem_InvalidArgument(t *testing.T) {
	testCases := []struct {
		name     string
		userId   int64
		itemType string
	}{
		{
			name:     "non existent item",
			userId:   1,
			itemType: "kettle",
		},
		{
			name:     "non existent user",
			userId:   50,
			itemType: "bipki",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var ctx = context.Background()
			fakeShop := fake.ShopServiceFake{
				FakeUsersById:   fakeUsersById,
				FakeItemsByType: fakeItemsByType,
			}
			fakeAuth := &fake.AuthServiceFake{Token: fakeToken}
			ctx = auth.SetUserIDToCtx(ctx, tc.userId)

			service := NewService(Deps{
				Shop: &fakeShop,
				Auth: fakeAuth,
			})

			_, err := service.BuyItem(ctx, &pb.BuyItemRequest{
				Item: tc.itemType,
			})
			if err == nil {
				t.Fatalf("expected error, got none")
			}
			if s, ok := status.FromError(err); !ok {
				t.Fatalf("no code in error")
			} else if s.Code() != codes.InvalidArgument {
				t.Fatalf("expected invalid argument code, got: %s", s.Code().String())
			}
		})
	}
}

func TestService_BuyItem_FailedPrecondition(t *testing.T) {
	var ctx = context.Background()
	fakeShop := fake.ShopServiceFake{
		FakeUsersById:   fakeUsersById,
		FakeItemsByType: fakeItemsByType,
	}
	fakeAuth := &fake.AuthServiceFake{Token: fakeToken}
	ctx = auth.SetUserIDToCtx(ctx, 1)

	service := NewService(Deps{
		Shop: &fakeShop,
		Auth: fakeAuth,
	})

	_, err := service.BuyItem(ctx, &pb.BuyItemRequest{
		Item: "bipki",
	})

	if err == nil {
		t.Fatalf("expected error, got none")
	}
	if s, ok := status.FromError(err); !ok {
		t.Fatalf("no code in error")
	} else if s.Code() != codes.FailedPrecondition {
		t.Fatalf("expected failed precondition code, got: %s", s.Code().String())
	}
}
