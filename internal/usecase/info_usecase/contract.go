package info_usecase

import (
	"context"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type infoStorage interface {
	GetWalletBalanceByUserID(ctx context.Context, userID uuid.UUID) (decimal.Decimal, error)
	GetInventoryByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Inventory, error)
	GetReceivedTransactionsByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Transaction, error)
	GetSentTransactionsByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Transaction, error)
}
