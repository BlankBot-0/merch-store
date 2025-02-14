package merch_platform

import "errors"

var (
	ErrNotEnoughCoins  = errors.New("not enough coins")
	ErrIncorrectAmount = errors.New("incorrect amount of coins")
)
