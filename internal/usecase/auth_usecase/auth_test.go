package auth_usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/auth_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/auth_usecase/mocks"
	"github.com/jackc/pgx/v5"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*gomock.Controller, *mocks.MockauthStorage, *mocks.MockjwtToken, *mocks.MocktrManager, *mocks.Mockpassword, *auth_usecase.AuthUsecase) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockauthStorage(ctrl)
	mockJWT := mocks.NewMockjwtToken(ctrl)
	mockTrManager := mocks.NewMocktrManager(ctrl)
	mockPassword := mocks.NewMockpassword(ctrl)

	u := auth_usecase.New(mockStorage, mockTrManager, mockJWT, mockPassword)

	return ctrl, mockStorage, mockJWT, mockTrManager, mockPassword, u
}

func TestAuth_SuccessfulLogin(t *testing.T) {
	ctrl, mockStorage, mockJWT, _, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	username := "testuser"
	password := "password123"
	passwordHash := "hashed_password"
	savedUser := domain.User{ID: userID, Username: username, PasswordHash: passwordHash}

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(savedUser, nil)
	mockPassword.EXPECT().ComparePassword(passwordHash, password).Return(nil)
	mockJWT.EXPECT().GenerateJWT(userID).Return("valid_token", nil)

	token, err := u.Auth(ctx, username, password)
	assert.NoError(t, err)
	assert.Equal(t, "valid_token", token)
}

func TestAuth_NewUserRegistration(t *testing.T) {
	ctrl, mockStorage, mockJWT, mockTrManager, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	username := "newuser"
	password := "newpassword"
	passwordHash := "hashed_password"
	userID := uuid.New()

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(domain.User{}, pgx.ErrNoRows)
	mockPassword.EXPECT().HashPassword(password).Return(passwordHash, nil)
	mockTrManager.EXPECT().Do(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	})
	mockStorage.EXPECT().CreateUser(ctx, gomock.Any()).Return(userID, nil)
	mockStorage.EXPECT().CreateWallet(ctx, gomock.Any()).Return(nil)
	mockJWT.EXPECT().GenerateJWT(userID).Return("valid_token", nil)

	token, err := u.Auth(ctx, username, password)
	assert.NoError(t, err)
	assert.Equal(t, "valid_token", token)
}

func TestAuth_InvalidPassword(t *testing.T) {
	ctrl, mockStorage, _, _, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	username := "testuser"
	password := "wrongpassword"
	passwordHash := "hashed_password"
	savedUser := domain.User{ID: uuid.New(), Username: username, PasswordHash: passwordHash}

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(savedUser, nil)
	mockPassword.EXPECT().ComparePassword(passwordHash, password).Return(errors.New("invalid password"))

	token, err := u.Auth(ctx, username, password)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuth_GenerateJWTError(t *testing.T) {
	ctrl, mockStorage, mockJWT, _, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	username := "testuser"
	password := "password123"
	passwordHash := "hashed_password"
	savedUser := domain.User{ID: userID, Username: username, PasswordHash: passwordHash}

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(savedUser, nil)
	mockPassword.EXPECT().ComparePassword(passwordHash, password).Return(nil)
	mockJWT.EXPECT().GenerateJWT(userID).Return("", errors.New("JWT generation failed"))

	token, err := u.Auth(ctx, username, password)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuth_HashPasswordError(t *testing.T) {
	ctrl, mockStorage, _, _, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	username := "newuser"
	password := "newpassword"

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(domain.User{}, pgx.ErrNoRows)
	mockPassword.EXPECT().HashPassword(password).Return("", errors.New("hashing failed"))

	token, err := u.Auth(ctx, username, password)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuth_CreateUserError(t *testing.T) {
	ctrl, mockStorage, _, mockTrManager, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	username := "newuser"
	password := "newpassword"
	passwordHash := "hashed_password"

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(domain.User{}, pgx.ErrNoRows)
	mockPassword.EXPECT().HashPassword(password).Return(passwordHash, nil)
	mockTrManager.EXPECT().Do(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	})
	mockStorage.EXPECT().CreateUser(ctx, gomock.Any()).Return(uuid.Nil, errors.New("failed to create user"))

	token, err := u.Auth(ctx, username, password)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuth_CreateWalletError(t *testing.T) {
	ctrl, mockStorage, _, mockTrManager, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	username := "newuser"
	password := "newpassword"
	passwordHash := "hashed_password"
	userID := uuid.New()

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(domain.User{}, pgx.ErrNoRows)
	mockPassword.EXPECT().HashPassword(password).Return(passwordHash, nil)
	mockTrManager.EXPECT().Do(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	})
	mockStorage.EXPECT().CreateUser(ctx, gomock.Any()).Return(userID, nil)
	mockStorage.EXPECT().CreateWallet(ctx, gomock.Any()).Return(errors.New("failed to create wallet"))

	token, err := u.Auth(ctx, username, password)
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestAuth_GetUserByUsernameError(t *testing.T) {
	ctrl, mockStorage, _, mockTrManager, mockPassword, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	username := "newuser"
	password := "newpassword"
	someError := errors.New("failed GetUserByUsername")

	mockStorage.EXPECT().GetUserByUsername(ctx, username).Return(domain.User{}, someError)
	mockPassword.EXPECT().HashPassword(gomock.Any()).Times(0)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).Times(0)
	mockStorage.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(0)
	mockStorage.EXPECT().CreateWallet(gomock.Any(), gomock.Any()).Times(0)

	token, err := u.Auth(ctx, username, password)
	assert.ErrorIs(t, err, someError)
	assert.Empty(t, token)
}
