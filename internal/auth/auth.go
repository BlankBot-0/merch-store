package auth

import (
	"Merch/internal/config"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Auth struct {
	PrivateKey     string
	ExpirationTime time.Duration
}

func New(cfg config.Auth) *Auth {
	return &Auth{
		PrivateKey:     cfg.PrivateKey,
		ExpirationTime: cfg.ExpirationTime,
	}
}

func (a *Auth) UserToken(ctx context.Context, userId int64) (string, error) {
	expirationTime := time.Now().Add(a.ExpirationTime)
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *Auth) UserAuth(ctx context.Context, token string) (int64, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(a.PrivateKey), nil
	})
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		return 0, ErrUnauthorized
	}
	if err != nil {
		return 0, ErrInvalidToken
	}
	if !tkn.Valid {
		return 0, ErrUnauthorized
	}
	return tkn.Claims.(jwt.MapClaims)["user_id"].(int64), nil
}

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}
