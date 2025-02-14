package postgres

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound = errors.New("entity not found")
)

func formatError(queryName string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("executing %s: %w", queryName, err)
}

func errIsNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
