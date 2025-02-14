package buyitem_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *storage) UpsertInventory(ctx context.Context, item string, userID uuid.UUID) error {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("inventory").
		Columns("user_id", "item", "quantity").
		Values(userID, item, 1).
		Suffix("ON CONFLICT (user_id, item) DO UPDATE SET quantity = inventory.quantity + 1, updated_at = CURRENT_TIMESTAMP").
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)
	_, err = conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
