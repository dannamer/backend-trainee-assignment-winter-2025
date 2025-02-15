package auth_usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (u *AuthUsecase) Auth(ctx context.Context, username, password string) (string, error) {
	user, err := u.storage.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			user.Username = username
			user.PasswordHash, err = u.password.HashPassword(password)
			if err != nil {
				return "", fmt.Errorf("error hashPassword: %w", err)
			}
			if err := u.trManager.Do(ctx, func(ctx context.Context) error {
				if user.ID, err = u.storage.CreateUser(ctx, user); err != nil {
					return fmt.Errorf("error CreateUser: %w", err)
				}
				wallet := domain.Wallet{
					UserID:  user.ID,
					Balance: 1000,
				}
				if err = u.storage.CreateWallet(ctx, wallet); err != nil {
					return fmt.Errorf("error CreateWallet: %w", err)
				}
				return nil
			}); err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("error GetUserByUsername: %w", err)
		}
	} else {
		if err = u.password.ComparePassword(user.PasswordHash, password); err != nil {
			return "", err
		}
	}

	token, err := u.jwt.GenerateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("error GenerateJWT: %w", err)
	}

	return token, nil
}
