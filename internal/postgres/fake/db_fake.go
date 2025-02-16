package fake

import (
	"Merch/internal/postgres"
	"context"
	"github.com/jackc/pgx/v5"
)

var _ postgres.DB = &DBMock{}

// DBMock is a type that implements DBInterface and aggregates repository mocks.
type DBMock struct {
	*UsersRepoFake
	*ShopRepoFake
}

func (d DBMock) ROShop() postgres.ROShop {
	return d.ShopRepoFake
}

func (d DBMock) RWShop() postgres.RWShop {
	return d.ShopRepoFake
}

func (d DBMock) ROUsers() postgres.ROUsers {
	return d.UsersRepoFake
}

func (d DBMock) RWUsers() postgres.RWUsers {
	return d.UsersRepoFake
}

func (d DBMock) RunInTx(_ context.Context, f func(tx postgres.RepositoryProvider) error, _ pgx.TxIsoLevel) error {
	return f(d)
}
