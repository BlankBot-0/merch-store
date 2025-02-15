package auth

import (
	"Merch/internal/auth"
	"Merch/internal/models"
	"context"
)

type (
	Authenticator interface {
		UserToken(ctx context.Context, userId int64) (string, error)
		UserAuth(ctx context.Context, token string) (int64, error)
	}
	CredentialsRepository interface {
		UserPasswordId(ctx context.Context, userLogin string) (models.InfoForTokenDTO, error)
	}
	RepositoryProvider interface {
		CredentialsRepository() CredentialsRepository
	}
)

type Deps struct {
	Authenticator Authenticator
	Repo          CredentialsRepository
}
type AuthSystem struct {
	Deps
}

func NewAuthenticationSystem(deps Deps) *AuthSystem {
	return &AuthSystem{
		Deps: deps,
	}
}

func (a *AuthSystem) UserToken(ctx context.Context, credentials models.CredentialsDTO) (string, error) {
	tokenInfo, err := a.Repo.UserPasswordId(ctx, credentials.Username)
	if err != nil {
		return "", auth.ErrUnauthorized
	}
	if tokenInfo.Password != credentials.Password {
		return "", auth.ErrUnauthorized
	}

	token, err := a.Deps.Authenticator.UserToken(ctx, tokenInfo.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthSystem) UserAuth(ctx context.Context, token string) (int64, error) {
	return a.Deps.Authenticator.UserAuth(ctx, token)
}
