package buyitem_usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
)

type buyItemStorage interface {
	GetWalletByUserID(ctx context.Context, userID uuid.UUID) (domain.Wallet, error)
	GetMerchByItem(ctx context.Context, item string) (domain.Merch, error)
	UpdateWallet(ctx context.Context, wallet domain.Wallet) error
	UpsertInventory(ctx context.Context, item string, userID uuid.UUID) error
}

type trManager interface {
	Do(context.Context, func(ctx context.Context) error) error
}
