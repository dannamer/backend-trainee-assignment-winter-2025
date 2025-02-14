package auth_usecase

import (
	"context"
	"fmt"
)

func (u *authUsecase) Auth(ctx context.Context, username, password string) (string, error) {
	user, err := u.storage.GetUserByUsername(ctx, username)
	if err != nil {
		passwordHash, err := u.password.HashPassword(password)
		if err != nil {
			return "", fmt.Errorf("error hashPassword: %w", err)
		}

		if err := u.trManager.Do(ctx, func(ctx context.Context) error {
			if user, err = u.storage.CreateUser(ctx, username, passwordHash); err != nil {
				return fmt.Errorf("error CreateUser: %w", err)
			}

			if err = u.storage.CreateWallet(ctx, user.ID); err != nil {
				return fmt.Errorf("error CreateWallet: %w", err)
			}
			return nil
		}); err != nil {
			return "", err
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
