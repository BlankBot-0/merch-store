package auth

import (
	"Merch/internal/postgres/fake"
	fi "Merch/internal/usecase/auth/fake"
	"context"
	"errors"
	"testing"
)

func TestAuthService_UserToken(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		name        string
		db          fake.DBMock
		issuer      Issuer
		expectedErr error
	}{
		{
			name: "ok",
			db: fake.DBMock{
				UsersRepoFake: &fake.UsersRepoFake{},
			},
			issuer: &fi.Issuer{
				Token: "token",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			auth := NewAuthService(Deps{
				Repo:   tc.db,
				Issuer: tc.issuer,
			})

			// issue token first time to register user
			token, err := auth.UserToken(ctx, "login", "password")
			if tc.expectedErr == nil && err != nil {
				t.Fatalf("got unexpected error: %s", err)
			} else if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("expected %s error, got %s", tc.expectedErr, err)
			}

			if token != "token" {
				t.Fatalf("got unexpected token `%s`, expected `token`", token)
			}

			// issue token one more time without registration
			token, err = auth.UserToken(ctx, "login", "password")
			if tc.expectedErr == nil && err != nil {
				t.Fatalf("got unexpected error: %s", err)
			} else if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("expected %s error, got %s", tc.expectedErr, err)
			}

			if token != "token" {
				t.Fatalf("got unexpected token `%s`, expected `token`", token)
			}
		})
	}
}
