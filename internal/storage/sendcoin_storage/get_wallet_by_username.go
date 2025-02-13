package sendcoin_storage

import (
	"context"
	std_errors "errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/jackc/pgx/v5"
)

func (s *storage) GetWalletByUsername(ctx context.Context, username string) (domain.Wallet, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("w.id", "w.user_id", "w.balance").
		From("wallet AS w").
		Join("users AS u ON u.id = w.user_id").
		Where(squirrel.Eq{"u.username": username}).
		ToSql()
	if err != nil {
		return domain.Wallet{}, fmt.Errorf("error sql build: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	var wallet domain.Wallet
	err = conn.QueryRow(ctx, query, args...).Scan(&wallet.ID, &wallet.UserID, &wallet.Balance)
	if err != nil {
		if std_errors.Is(err, pgx.ErrNoRows) {
			return domain.Wallet{}, errors.ErrUserNotFound
		}
		return domain.Wallet{}, fmt.Errorf("error QueryRow: %w", err)
	}

	return wallet, nil
}
