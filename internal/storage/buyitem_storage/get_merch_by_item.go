package buyitem_storage

import (
	"context"
	std_errors "errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/jackc/pgx/v5"
)

func (s *storage) GetMerchByItem(ctx context.Context, item string) (domain.Merch, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("id", "item", "price").
		From("merch_store").
		Where(squirrel.Eq{"item": item}).
		ToSql()
	if err != nil {
		return domain.Merch{}, fmt.Errorf("error sql build: %w", err)
	}

	var merch domain.Merch
	err = s.pg.QueryRow(ctx, query, args...).Scan(&merch.ID, &merch.Item, &merch.Price)
	if err != nil {
		if std_errors.Is(err, pgx.ErrNoRows) {
			return domain.Merch{}, errors.ErrMerchNotFound
		}
		return domain.Merch{}, fmt.Errorf("error QueryRow: %w", err)
	}

	return merch, nil
}
