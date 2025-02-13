package info_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
)

func (s *storage) GetInventoryByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Inventory, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("item", "quantity").
		From("inventory").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error sql build: %w", err)
	}

	var inventory []domain.Inventory
	rows, err := s.pg.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var inv domain.Inventory
		if err := rows.Scan(&inv.Item, &inv.Quantity); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		inventory = append(inventory, inv)
	}

	return inventory, nil
}
