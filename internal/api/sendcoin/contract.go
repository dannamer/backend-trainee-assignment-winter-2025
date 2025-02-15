package sendcoin

import (
	"context"

	"github.com/google/uuid"
)

type sendcoinUsecase interface {
	SendCoin(ctx context.Context, toUsername string, userID uuid.UUID, amout int64) error
}
