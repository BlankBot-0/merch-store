package auth

import (
	"Merch/internal/auth"
	"Merch/internal/models"
	"Merch/internal/postgres"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type (
	Issuer interface {
		Issue(userId int64) (string, error)
	}
)

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

func (a *AuthService) UserToken(ctx context.Context, credentials models.CredentialsDTO) (string, error) {
	passwordHashRaw, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}

	passwordHash := string(passwordHashRaw)
	tokenInfo, err := a.Repo.ROUsers().UserByLogin(ctx, credentials.Username)

	switch {
	case errors.Is(err, postgres.ErrNotFound):
		id, err := a.createUser(ctx, credentials.Username, passwordHash)
		if err != nil {
			return "", fmt.Errorf("could not create user: %w", err)
		}
		tokenInfo.Id = id
	case err != nil:
		return "", fmt.Errorf("could not get user: %w", err)
	default:
		if tokenInfo.PasswordHash != passwordHash {
			return "", auth.ErrUnauthorized
		}
	}

	token, err := a.Deps.Authenticator.Issue(tokenInfo.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) createUser(ctx context.Context, login, passwordHash string) (int64, error) {
	id, err := a.Deps.Repo.RWUsers().CreateUser(ctx, login, passwordHash)
	if errors.Is(err, postgres.ErrAlreadyExists) {
		return 0, ErrUserAlreadyExists
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

var ErrUserAlreadyExists = errors.New("user already exists")
