package auth_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *storage) CreateWallet(ctx context.Context, ID uuid.UUID) error {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("wallet").
		Columns("user_id").
		Values(ID).
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
