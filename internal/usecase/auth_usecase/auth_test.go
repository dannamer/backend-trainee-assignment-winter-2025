package auth_usecase_test

import (
	"context"
	"testing"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/auth_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/auth_usecase/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthUsecase_Auth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockauthStorage(ctrl)
	mockJWT := mocks.NewMockjwtToken(ctrl)
	mockTrManager := mocks.NewMocktrManager(ctrl)
	mockPassword := mocks.NewMockpassword(ctrl)

	u := auth_usecase.New(mockStorage, mockTrManager, mockJWT, mockPassword)

	ctx := context.Background()
	userID := uuid.New()
	username := "testuser"
	password := "password123"
	passwordHash := "hashed_password"
	savedUser := domain.User{ID: userID, Username: username, PasswordHash: passwordHash}

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(savedUser, nil)
	mockJWT.EXPECT().GenerateJWT(userID).Return("valid_token", nil)

	token, err := u.Auth(ctx, username, password)
	assert.NoError(t, err)
	assert.Equal(t, "valid_token", token)
}
