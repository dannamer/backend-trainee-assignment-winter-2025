package sendcoin_usecase

import (
	"context"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
)

type sendCoinStorage interface {
	GetWalletByUserID(ctx context.Context, userID uuid.UUID) (domain.Wallet, error)
	GetWalletByUsername(ctx context.Context, username string) (domain.Wallet, error)
	UpdateWallet(ctx context.Context, wallet domain.Wallet) error
	CreateTransactions(ctx context.Context, senderID, receiverID uuid.UUID, amount int) error
}

type trManager interface {
	Do(context.Context, func(ctx context.Context) error) error
}
