package fake

import (
	"Merch/internal/models"
	"Merch/internal/usecase/shop"
	"context"
)

type ShopServiceFake struct {
	FakeInfo         models.UserInfo
	FakeUsersById    map[int64]UserFake
	FakeUsersByLogin map[string]UserFake
}

type UserFake struct {
	Id    int64
	Login string
	Coins int64
}

func (s *ShopServiceFake) BuyItem(_ context.Context, _ int64, _ string) error {
	return nil
}

func (s *ShopServiceFake) SendCoins(_ context.Context, receiverId int64, _ string, _ int64) error {
	if _, ok := s.FakeUsersById[receiverId]; !ok {
		return shop.ErrUserIsNotFound
	}
	return nil
}

func (s *ShopServiceFake) Info(_ context.Context, _ int64) (models.UserInfo, error) {
	return s.FakeInfo, nil
}
