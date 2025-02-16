package auth

import (
	"Merch/internal/config"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims interface {
	UserID() int64
}

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

func (a *Auth) Issue(userId int64) (string, error) {
	expirationTime := time.Now().Add(a.ExpirationTime)
	claims := &privateClaims{
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

func (a *Auth) Verify(token string) (Claims, error) {
	claims := &privateClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(a.PrivateKey), nil
	})
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		return nil, ErrUnauthorized
	}
	if err != nil {
		return nil, ErrInvalidToken
	}
	if !tkn.Valid {
		return nil, ErrUnauthorized
	}
	return claims, nil
}

type privateClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (p *privateClaims) UserID() int64 {
	return p.UserId
}

type userIDKey struct{}

func GetUserIDFromCtx(ctx context.Context) int64 {
	userID, ok := ctx.Value(userIDKey{}).(int64)
	if !ok {
		return 0
	}
	return userID
}

func SetUserIDToCtx(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}
