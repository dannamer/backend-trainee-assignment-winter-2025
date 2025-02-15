package info_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *storage) GetWalletBalanceByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("balance").
		From("wallet").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("error sql build: %w", err)
	}

	var balance int64
	err = s.pg.QueryRow(ctx, query, args...).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("error query row: %w", err)
	}

	return balance, nil
}
