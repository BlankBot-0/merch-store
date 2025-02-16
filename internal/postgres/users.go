package postgres

import (
	"Merch/internal/models"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type roUsers struct {
	query querier
}

func (ro *roUsers) UserById(ctx context.Context, userId int64) (models.User, error) {
	const queryName = "UsersRepository/UserCoins"
	const q = `
        select id, login, password_hash, coins
        from users
        where id = $1`

	var user models.User
	if err := pgxscan.Get(ctx, ro.query, &user, q, userId); errIsNoRows(err) {
		return user, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return user, formatError(queryName, err)
	}

	return user, nil
}

func (ro *roUsers) UserByLogin(ctx context.Context, login string) (models.User, error) {
	const queryName = "UsersRepository/UserByLogin"
	const q = `
        select id, login, password_hash, coins
        from users
        where login = $1`

	var user models.User
	if err := pgxscan.Get(ctx, ro.query, &user, q, login); errIsNoRows(err) {
		return user, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return user, formatError(queryName, err)
	}

	return user, nil
}

type rwUsers struct {
	ROUsers
	exec executor
}

func (rw *rwUsers) CreateUser(ctx context.Context, login, passwordHash string) (int64, error) {
	const queryName = "UsersRepository/createUser"
	const q = `
		insert into users(login, password_hash, coins)
		values ($1, $2, 1000)
		returning id`

	var id int64
	if err := pgxscan.Get(ctx, rw.exec, &id, q, login, passwordHash); isUniqueViolated(err) {
		return 0, formatError(queryName, ErrAlreadyExists)
	} else if err != nil {
		return 0, formatError(queryName, err)
	}

	return id, nil
}
