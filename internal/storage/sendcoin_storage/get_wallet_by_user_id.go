package sendcoin_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
)

func (s *storage) GetWalletByUserID(ctx context.Context, userID uuid.UUID) (domain.Wallet, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("id", "user_id", "balance").
		From("wallet").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return domain.Wallet{}, fmt.Errorf("error sql build: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	var wallet domain.Wallet
	err = conn.QueryRow(ctx, query, args...).Scan(&wallet.ID, &wallet.UserID, &wallet.Balance)
	if err != nil {
		return domain.Wallet{}, fmt.Errorf("error QueryRow: %w", err)
	}

	return wallet, nil
}
