package buyitem

import (
	"context"

	"github.com/google/uuid"
)

type buyItemUsecase interface {
	BuyItem(ctx context.Context, userID uuid.UUID, item string) error
}
