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
	FakeItemsByType  map[string]models.Item
}

type UserFake struct {
	Id    int64
	Login string
	Coins int64
}

func (s *ShopServiceFake) BuyItem(_ context.Context, userId int64, itemType string) error {
	usr, ok := s.FakeUsersById[userId]
	if !ok {
		return shop.ErrUserIsNotFound
	}
	itm, ok := s.FakeItemsByType[itemType]
	if !ok {
		return shop.ErrItemIsNotFound
	}
	if usr.Coins < itm.Coins {
		return shop.ErrNotEnoughCoins
	}
	return nil
}

func (s *ShopServiceFake) SendCoins(_ context.Context, senderId int64, receiverLogin string, amount int64) error {
	if _, ok := s.FakeUsersByLogin[receiverLogin]; !ok {
		return shop.ErrUserIsNotFound
	}
	if sender, ok := s.FakeUsersById[senderId]; !ok {
		return shop.ErrUserIsNotFound
	} else if sender.Coins < amount {
		return shop.ErrNotEnoughCoins
	}
	return nil
}

func (s *ShopServiceFake) Info(_ context.Context, userid int64) (*models.UserInfo, error) {
	return &s.FakeInfo, nil
}
