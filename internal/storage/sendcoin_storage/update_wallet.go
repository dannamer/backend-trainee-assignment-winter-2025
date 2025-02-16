package sendcoin_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
)

func (s *storage) UpdateWallet(ctx context.Context, wallet domain.Wallet) error {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update("wallet").
		Set("balance", wallet.Balance).
		Set("updated_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Where(squirrel.Eq{"id": wallet.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("error sql build: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	_, err = conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update wallet: %w", err)
	}

	return nil
}
