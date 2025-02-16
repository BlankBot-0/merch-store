package fake

import (
	"Merch/internal/models"
	"Merch/internal/postgres"
	"context"
	"github.com/samber/lo"
)

var _ postgres.RWUsers = &UsersRepoFake{}

type UsersRepoFake struct {
	lastIdx int64
	users   []models.User

	CreateUserErr  error
	UserByIdErr    error
	UserByLoginErr error
}

func (u *UsersRepoFake) CreateUser(_ context.Context, login, passwordHash string) (int64, error) {
	if u.CreateUserErr != nil {
		return 0, u.CreateUserErr
	}
	u.lastIdx++
	u.users = append(u.users, models.User{
		Id:           u.lastIdx,
		Login:        login,
		PasswordHash: passwordHash,
		Coins:        1000,
	})
	return u.lastIdx, nil
}

func (u *UsersRepoFake) UserById(_ context.Context, userId int64) (models.User, error) {
	if u.UserByIdErr != nil {
		return models.User{}, u.UserByIdErr
	}

	user, ok := lo.Find(u.users, func(item models.User) bool {
		return item.Id == userId
	})
	if !ok {
		return user, postgres.ErrNotFound
	}

	return user, nil
}

func (u *UsersRepoFake) UserByLogin(_ context.Context, login string) (models.User, error) {
	if u.UserByLoginErr != nil {
		return models.User{}, u.UserByLoginErr
	}

	user, ok := lo.Find(u.users, func(item models.User) bool {
		return item.Login == login
	})
	if !ok {
		return user, postgres.ErrNotFound
	}

	return user, nil
}
