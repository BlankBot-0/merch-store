package postgres

import (
	"Merch/internal/usecase/merch_platform"
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func Connect(ctx context.Context, dsn string) (*Database, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("initializing postgres: %w", err)
	}

	return &Database{pool: pool}, nil
}

func (db *Database) Close() {
	db.pool.Close()
}

// ROSalesPlatform returns a user postgres which will use this database for querying.
func (db *Database) ROSalesPlatform() merch_platform.ROSalesPlatform {
	return &roSalesPlatformRepository{query: dbReader{db.pool}}
}

// RWSalesPlatform returns a user postgres which will use this database for querying and executing.
func (db *Database) RWSalesPlatform() merch_platform.RWSalesPlatform {
	return &rwSalesPlatformRepository{
		ROSalesPlatform: db.ROSalesPlatform(),
		exec:            dbWriter{db.pool},
	}
}

// WriteTx is an active writeable and readable transaction launched by a Database instance.
// Repository methods accessed through WriteTx are run in this transaction.
type WriteTx struct {
	wrapped pgx.Tx
}

// ROSalesPlatform returns a user postgres which will user this transaction for querying.
func (tx *WriteTx) ROSalesPlatform() merch_platform.ROSalesPlatform {
	return &roSalesPlatformRepository{query: tx.wrapped}
}

// RWSalesPlatform returns a user postgres which will user this transaction for querying and execution.
func (tx *WriteTx) RWSalesPlatform() merch_platform.RWSalesPlatform {
	return &rwSalesPlatformRepository{
		ROSalesPlatform: tx.ROSalesPlatform(),
		exec:            tx.wrapped,
	}
}

// RunInTx runs the specified function in a transaction which supports writing and reading.
func (db *Database) RunInTx(ctx context.Context, f func(tx merch_platform.RepositoryProvider) error, isoLevel pgx.TxIsoLevel) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: isoLevel})
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}
	return f(&WriteTx{wrapped: tx})
}

// querier can only be used for reading data
type querier = pgxscan.Querier

// executor can be used both for reading and writing data
type executor interface {
	querier
	Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error)
}

// dbReader implements querier with a read-only database connection.
type dbReader struct {
	pool *pgxpool.Pool
}

func (cr dbReader) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	rows, err := cr.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	} else if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return rows, nil
}

// dbWriter implements executor with a read-write cluster connection.
type dbWriter struct {
	pool *pgxpool.Pool
}

func (cw dbWriter) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	rows, err := cw.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	} else if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return rows, nil
}

func (cw dbWriter) Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error) {
	return cw.pool.Exec(ctx, sql, args...)
}
