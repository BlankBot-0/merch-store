package postgres

import (
	"Merch/internal/models"
	"Merch/internal/usecase/merch_platform"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type roSalesPlatformRepository struct {
	query querier
}

func (ro *roSalesPlatformRepository) UserById(ctx context.Context, userId int64) (models.User, error) {
	const queryName = "SalesPlatformRepository/UserCoins"
	const q = `
               select login, coins from users where id = $1`

	var user models.User
	if err := pgxscan.Get(ctx, ro.query, &user, q, userId); errIsNoRows(err) {
		return user, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return user, formatError(queryName, err)
	}

	return user, nil
}

func (ro *roSalesPlatformRepository) UserByLogin(ctx context.Context, login string) (models.User, error) {
	const queryName = "SalesPlatformRepository/UserByLogin"
	const q = `
               select id, login, coins from users where login = $1`

	var user models.User
	if err := pgxscan.Get(ctx, ro.query, &user, q, login); errIsNoRows(err) {
		return user, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return user, formatError(queryName, err)
	}

	return user, nil
}

func (ro *roSalesPlatformRepository) Item(ctx context.Context, itemType string) (models.Item, error) {
	const queryName = "SalesPlatformRepository/Item"
	const q = `
               select id, type, quantity from merchandise where type = $1`

	var item models.Item

	if err := pgxscan.Get(ctx, ro.query, &item, q, itemType); errIsNoRows(err) {
		return item, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return item, formatError(queryName, err)
	}

	return item, nil
}

func (ro *roSalesPlatformRepository) UserInventory(ctx context.Context, userId int64) ([]models.InventoryItem, error) {
	const queryName = "SalesPlatformRepository/UserInventory"
	const q = `
               select merchandise.type, count(*) as quantity
               from users
               join purchases on users.id = purchases.user_id
               join merchandise on purchases.item_id = merchandise.id
               where users.id = $1
               group by purchases.id`

	var items []models.InventoryItem
	if err := pgxscan.Get(ctx, ro.query, &items, q, userId); errIsNoRows(err) {
		return items, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return items, formatError(queryName, err)
	}

	return items, nil
}

func (ro *roSalesPlatformRepository) UserSendOperations(ctx context.Context, userId int64) ([]models.Transaction, error) {
	const queryName = "SalesPlatformRepository/UserSendOperations"
	const q = `
               select users.login, sent.amount
               from (
                   select to_user_id, amount
                   from transactions
                   where from_user_id = $1
               ) sent
               join users on sent.to_user_id = users.id`

	var sent []models.Transaction
	if err := pgxscan.Get(ctx, ro.query, &sent, q, userId); errIsNoRows(err) {
		return sent, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return sent, formatError(queryName, err)
	}

	return sent, nil
}

func (ro *roSalesPlatformRepository) UserReceivedOperations(ctx context.Context, userId int64) ([]models.Transaction, error) {
	const queryName = "SalesPlatformRepository/UserReceivedOperations"
	const q = `
               select users.login, sent.amount
               from (
                   select from_user_id, amount
                   from transactions
                   where to_user_id = $1
               ) sent
               join users on sent.from_user_id = users.id`

	var received []models.Transaction
	if err := pgxscan.Get(ctx, ro.query, &received, q, userId); errIsNoRows(err) {
		return received, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return received, formatError(queryName, err)
	}

	return received, nil
}

func (ro *roSalesPlatformRepository) UserPasswordId(ctx context.Context, userLogin string) (models.InfoForTokenDTO, error) {
	const queryName = "SalesPlatformRepository/UserPassword"
	const q = `
               select login, id from users where login = $1`
	var cred models.InfoForTokenDTO
	if err := pgxscan.Get(ctx, ro.query, &cred, q, userLogin); err != nil {
		return cred, formatError(queryName, err)
	}
	return cred, nil
}

type rwSalesPlatformRepository struct {
	merch_platform.ROSalesPlatform
	exec executor
}

func (rw *rwSalesPlatformRepository) RemoveCoins(ctx context.Context, userId, amount int64) error {
	const queryName = "SalesPlatformRepository/SendCoins"
	const q = `
               update users set coins = coins - $2 where id = $1 returng coins`

	var coins int64
	if err := pgxscan.Get(ctx, rw.exec, &coins, q, userId, amount); err != nil {
		return formatError(queryName, err)
	}

	return nil
}

func (rw *rwSalesPlatformRepository) AddCoins(ctx context.Context, userId, amount int64) error {
	const queryName = "SalesPlatformRepository/AddCoins"
	const q = `
               update users set coins = coins + $2 where id = $1`

	if err := pgxscan.Get(ctx, rw.exec, nil, q, userId, amount); err != nil {
		return formatError(queryName, err)
	}

	return nil
}

func (rw *rwSalesPlatformRepository) AddTransaction(ctx context.Context, fromUserId, toUserId int64, amount int64) error {
	const queryName = "SalesPlatformRepository/AddTransaction"
	const q = `
               insert into transactions(from_user_id, to_user_id, amount)
               values ($1, $2, $3)`

	if err := pgxscan.Get(ctx, rw.exec, nil, q, fromUserId, toUserId, amount); err != nil {
		return formatError(queryName, err)
	}

	return nil
}

func (rw *rwSalesPlatformRepository) AddPurchase(ctx context.Context, userId, itemId int64) error {
	const queryName = "SalesPlatformRepository/Purchase"
	const q = `
               insert into purchases(user_id, item_id)
               values ($1, $2)`

	if err := pgxscan.Get(ctx, rw.exec, nil, q, userId, itemId); errIsNoRows(err) {
		return formatError(queryName, ErrNotFound)
	} else if err != nil {
		return formatError(queryName, err)
	}

	return nil
}
