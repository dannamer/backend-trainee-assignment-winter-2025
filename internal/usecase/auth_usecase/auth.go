package auth_usecase

import (
	"context"
	"fmt"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *authUsecase) Auth(ctx context.Context, username, password string) (string, error) {
	user, err := u.storage.GetUserByUsername(ctx, username)
	if err != nil {
		passwordHash, err := hashPassword(password)
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
	}

	if err = comparePassword(user.PasswordHash, password); err != nil {
		return "", errors.ErrInvalidPassword
	}

	token, err := u.jwt.GenerateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("error GenerateJWT: %w", err)
	}

	return token, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
