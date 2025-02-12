package buyitem_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *storage) SavePurchase(ctx context.Context, userID, merchID uuid.UUID) error {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("purchase_history").
		Columns("user_id", "merch_id").
		Values(userID, merchID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error sql build user: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	_, err = conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert purchase history: %w", err)
	}

	return nil
}
