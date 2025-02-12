package auth_storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
)

func (s *storage) CreateUser(ctx context.Context, username, passwordHash string) (domain.User, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("users").
		Columns("username", "password").
		Values(username, passwordHash).
		Suffix("RETURNING id, username, password").
		ToSql()
	if err != nil {
		return domain.User{}, fmt.Errorf("error sql build user: %w", err)
	}
	
	conn := s.trGetter.DefaultTrOrDB(ctx, s.pg)

	var user domain.User
	err = conn.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return domain.User{}, fmt.Errorf("error QueryRow user: %w", err)
	}

	return user, nil
}