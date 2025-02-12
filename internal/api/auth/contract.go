package auth

import (
	"context"
)

type authUsecase interface {
	Auth(ctx context.Context, username, password string) (string, error)
}