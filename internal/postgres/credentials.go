package postgres

import (
	"Merch/internal/models"
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type credentialsRepository struct {
	query querier
}

func (c *credentialsRepository) UserPasswordId(ctx context.Context, userLogin string) (models.InfoForTokenDTO, error) {
	const queryName = "CredentialsRepository/UserPasswordId"
	const q = `
               select password, id from users where login = $1`
	var cred models.InfoForTokenDTO
	if err := pgxscan.Get(ctx, c.query, &cred, q, userLogin); err != nil {
		return cred, formatError(queryName, err)
	}
	return cred, nil
}
