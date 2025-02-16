package fake

import (
	"context"
)

type AuthServiceFake struct {
	Token string
	Err   error
}

func (a AuthServiceFake) UserToken(_ context.Context, _, _ string) (string, error) {
	if a.Err != nil {
		return "", a.Err
	}

	return a.Token, nil
}
