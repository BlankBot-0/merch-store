package shop

import "errors"

var (
	ErrNotEnoughCoins = errors.New("not enough coins")
	ErrUserIsNotFound = errors.New("user is not found")
	ErrItemIsNotFound = errors.New("item is not found")
)
