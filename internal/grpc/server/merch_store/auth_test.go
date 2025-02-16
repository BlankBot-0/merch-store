package merch_store

import (
	"Merch/internal/grpc/server/merch_store/fake"
	pb "Merch/pkg/api/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const (
	fakeToken = "token"
)

func TestService_Auth_Ok(t *testing.T) {
	fakeAuth := &fake.AuthServiceFake{
		Token: fakeToken,
	}

	service := NewService(Deps{
		Auth: fakeAuth,
	})

	resp, err := service.Auth(context.Background(), &pb.AuthRequest{
		Login:    "user",
		Password: "password",
	})
	if err != nil {
		t.Fatalf("got unexpected error: %s", err)
	}

	if resp.Token != fakeToken {
		t.Fatalf("unexpected token in response, expected: %s, got: %s", fakeToken, resp.Token)
	}
}

func TestService_Auth_Validation(t *testing.T) {
	testCases := []struct {
		name     string
		login    string
		password string
	}{
		{
			name:     "short login",
			login:    "l",
			password: "password",
		},
		{
			name:     "short password",
			login:    "login",
			password: "pwd",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeAuth := &fake.AuthServiceFake{Token: fakeToken}

			service := NewService(Deps{
				Auth: fakeAuth,
			})

			_, err := service.Auth(context.Background(), &pb.AuthRequest{
				Login:    tc.login,
				Password: tc.password,
			})
			if err == nil {
				t.Fatalf("expected validation error, but got no error")
			}

			if s, ok := status.FromError(err); !ok {
				t.Fatalf("no code in error")
			} else if s.Code() != codes.InvalidArgument {
				t.Fatalf("expected invalid argument code, got: %s", s.Code().String())
			}
		})
	}
}
