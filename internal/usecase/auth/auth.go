package auth

import (
	"Merch/internal/auth"
	"Merch/internal/postgres"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Issuer interface {
	Issue(userId int64) (string, error)
}

type Deps struct {
	Authenticator Issuer
	Repo          postgres.DB
}
type AuthService struct {
	Deps
}

func NewAuthService(deps Deps) *AuthService {
	return &AuthService{
		Deps: deps,
	}
}

func (a *AuthService) UserToken(ctx context.Context, login, password string) (string, error) {
	tokenInfo, err := a.Repo.ROUsers().UserByLogin(ctx, login)

	switch {
	case errors.Is(err, postgres.ErrNotFound):
		id, err := a.createUser(ctx, login, password)
		if err != nil {
			return "", fmt.Errorf("could not create user: %w", err)
		}
		tokenInfo.Id = id
	case err != nil:
		return "", fmt.Errorf("could not get user: %w", err)
	default:
		err := bcrypt.CompareHashAndPassword([]byte(tokenInfo.PasswordHash), []byte(password))
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", fmt.Errorf("incorrect password: %w", auth.ErrUnauthorized)
		} else if err != nil {
			return "", fmt.Errorf("failed to verify password: %w", err)
		}
	}

	token, err := a.Deps.Authenticator.Issue(tokenInfo.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) createUser(ctx context.Context, login, password string) (int64, error) {
	passwordHashRaw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return 0, fmt.Errorf("could not hash password: %w", err)
	}

	id, err := a.Deps.Repo.RWUsers().CreateUser(ctx, login, string(passwordHashRaw))
	if errors.Is(err, postgres.ErrAlreadyExists) {
		return 0, ErrUserAlreadyExists
	} else if err != nil {
		return 0, err
	}

	return id, nil
}
