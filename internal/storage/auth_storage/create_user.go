package auth_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
)

func (s *storage) CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("users").
		Columns("username", "password").
		Values(user.Username, user.PasswordHash).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error sql build user: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	var userID uuid.UUID
	err = conn.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error QueryRow user: %w", err)
	}

	return userID, nil
}
