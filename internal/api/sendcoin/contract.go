package sendcoin

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type sendcoinUsecase interface {
	SendCoin(ctx context.Context, toUsername string, userID uuid.UUID, amout decimal.Decimal) error
}
