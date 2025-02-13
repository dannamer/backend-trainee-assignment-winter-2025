package info_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
)

func (s *storage) GetReceivedTransactionsByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Transaction, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("u.username, t.amount").
		From("transactions t").
		Join("users u ON t.sender_id = u.id").
		Where(squirrel.Eq{"t.receiver_id": userID}).
		OrderBy("t.created_at DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error sql build: %w", err)
	}

	var transactions []domain.Transaction
	rows, err := s.pg.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Transaction
		if err := rows.Scan(&t.Username, &t.Amount); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return transactions, nil
}
