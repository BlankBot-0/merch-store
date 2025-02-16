package merch_store

import (
	"Merch/internal/grpc/server/merch_store/fake"
	pb "Merch/pkg/api/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var (
	fakeUsersById = map[int64]fake.UserFake{
		1: {1, "petya", 500},
		2: {2, "vasya", 300},
	}
	fakeUsersByLogin = map[string]fake.UserFake{
		"petya": {1, "petya", 500},
		"vasya": {2, "vasya", 300},
	}
)

func TestService_SendCoin_InvalidArgument(t *testing.T) {
	testCases := []struct {
		name     string
		receiver string
		amount   int64
	}{
		{
			name:     "negative amount",
			receiver: "petya",
			amount:   -10,
		},
		{
			name:     "non positive amount",
			receiver: "vasya",
			amount:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeShop := &fake.ShopServiceFake{FakeInfo: fakeInfo, FakeUsersById: fakeUsersById}

			service := NewService(Deps{
				Shop: fakeShop,
			})

			_, err := service.SendCoin(context.Background(), &pb.SendCoinRequest{
				ToUser: tc.receiver,
				Amount: tc.amount,
			})
			if err == nil {
				t.Fatal("expected error, got none")
			}

			if s, ok := status.FromError(err); !ok {
				t.Fatalf("no code error")
			} else if s.Code() != codes.InvalidArgument {
				t.Errorf("expected invalid code, got: %d", s.Code())
			}
		})
	}
}

func TestService_SendCoin



{
name:     "short login",
receiver: "",
amount:   40,
},