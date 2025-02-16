package postgres

import (
	"Merch/internal/models"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type roShop struct {
	query querier
}

func (ro *roShop) GetItem(ctx context.Context, itemType string) (models.Item, error) {
	const queryName = "SalesPlatformRepository/GetItem"
	const q = `
        select id, type, coins
        from items
        where type = $1`

	var item models.Item

	if err := pgxscan.Get(ctx, ro.query, &item, q, itemType); errIsNoRows(err) {
		return item, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return item, formatError(queryName, err)
	}

	return item, nil
}

func (ro *roShop) UserInventory(ctx context.Context, userId int64) ([]models.InventoryItem, error) {
	const queryName = "SalesPlatformRepository/UserInventory"
	const q = `
        select items.type, count(*) as quantity
        from users
        join purchases on users.id = purchases.user_id
        join items on purchases.item_id = items.id
        where users.id = $1
        group by items.type`

	var items []models.InventoryItem
	if err := pgxscan.Select(ctx, ro.query, &items, q, userId); errIsNoRows(err) {
		return items, formatError(queryName, ErrNotFound)
	} else if err != nil {
		return items, formatError(queryName, err)
	}

	return items, nil
}

func (ro *roShop) UserTransactions(ctx context.Context, userId int64) ([]models.Transaction, error) {
	const queryName = "SalesPlatformRepository/UserTransactions"
	const q = `
        select receiver.login as receiver, sender.login as sender, t.coins as amount
        from transactions t
        join users sender on t.from_user_id = sender.id
        join users receiver on t.to_user_id = receiver.id
        where sender.id = $1 or receiver.id = $1`

	var tx []models.Transaction
	if err := pgxscan.Select(ctx, ro.query, &tx, q, userId); err != nil {
		return tx, formatError(queryName, err)
	}

	return tx, nil
}

type rwShop struct {
	ROShop
	exec executor
}

func (rw *rwShop) AddCoins(ctx context.Context, userId, amount int64) error {
	const queryName = "SalesPlatformRepository/AddCoins"
	const q = `
        update users
        set coins = coins + $2
        where id = $1`

	if tag, err := rw.exec.Exec(ctx, q, userId, amount); err != nil {
		return formatError(queryName, err)
	} else if tag.RowsAffected() < 1 {
		return formatError(queryName, ErrNotChanged)
	}

	return nil
}

func (rw *rwShop) SaveTransaction(ctx context.Context, fromUserId, toUserId int64, amount int64) error {
	const queryName = "SalesPlatformRepository/SaveTransaction"
	const q = `
        insert into transactions(from_user_id, to_user_id, coins)
        values ($1, $2, $3)`

	if tag, err := rw.exec.Exec(ctx, q, fromUserId, toUserId, amount); err != nil {
		return formatError(queryName, err)
	} else if tag.RowsAffected() < 1 {
		return formatError(queryName, ErrNotChanged)
	}

	return nil
}

func (rw *rwShop) AddPurchase(ctx context.Context, userId, itemId int64) error {
	const queryName = "SalesPlatformRepository/Purchase"
	const q = `
        insert into purchases(user_id, item_id)
        values ($1, $2)`

	if tag, err := rw.exec.Exec(ctx, q, userId, itemId); err != nil {
		return formatError(queryName, err)
	} else if tag.RowsAffected() < 1 {
		return formatError(queryName, ErrNotChanged)
	}

	return nil
}
