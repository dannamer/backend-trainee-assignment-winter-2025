package auth_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
)

func (s *storage) CreateWallet(ctx context.Context, wallet domain.Wallet) error {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("wallet").
		Columns("user_id", "balance").
		Values(wallet.ID, wallet.Balance).
		ToSql()
	if err != nil {
		return fmt.Errorf("error sql build user: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	if _, err = conn.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("error Exec wallet: %w", err)
	}

	return nil
}
