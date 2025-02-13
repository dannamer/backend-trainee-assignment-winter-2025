package info_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (s *storage) GetWalletBalanceByUserID(ctx context.Context, userID uuid.UUID) (decimal.Decimal, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("balance").
		From("wallet").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("error sql build: %w", err)
	}

	var balance decimal.Decimal
	err = s.pg.QueryRow(ctx, query, args...).Scan(&balance)
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("error query row: %w", err)
	}

	return balance, nil
}
