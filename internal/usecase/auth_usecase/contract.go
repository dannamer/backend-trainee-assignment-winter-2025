package auth_usecase

import (
	"context"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
)

type authStorage interface {
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error)
	CreateWallet(ctx context.Context, wallet domain.Wallet) error
}

type jwtToken interface {
	GenerateJWT(ID uuid.UUID) (string, error)
}

type trManager interface {
	Do(context.Context, func(ctx context.Context) error) error
}

type password interface {
	ComparePassword(hashedPassword, password string) error
	HashPassword(password string) (string, error)
}
