package auth_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
)

func (s *storage) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("id", "username", "password").
		From("users").
		Where(squirrel.Eq{"username": username}).
		ToSql()
	if err != nil {
		return domain.User{}, fmt.Errorf("error sql build: %w", err)
	}

	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	var user domain.User
	err = conn.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return domain.User{}, fmt.Errorf("error QueryRow: %w", err)
	}

	return user, nil
}
