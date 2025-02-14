package sendcoin_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *storage) CreateTransactions(ctx context.Context, senderID, receiverID uuid.UUID, amount decimal.Decimal) error {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("transactions").
		Columns("sender_id", "receiver_id", "amount").
		Values(senderID, receiverID, amount).
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
