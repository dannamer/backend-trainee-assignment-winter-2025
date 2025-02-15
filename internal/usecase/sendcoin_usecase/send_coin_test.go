package sendcoin_usecase_test

import (
	"context"
	std_errors "errors"
	"testing"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/sendcoin_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/sendcoin_usecase/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*gomock.Controller, *mocks.MocksendCoinStorage, *mocks.MocktrManager, *sendcoin_usecase.SendCoinUsecase) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMocksendCoinStorage(ctrl)
	mockTrManager := mocks.NewMocktrManager(ctrl)
	u := sendcoin_usecase.New(mockStorage, mockTrManager)
	return ctrl, mockStorage, mockTrManager, u
}

func TestSendCoin_Success(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	toUserID := uuid.New()
	toUsername := "receiver"
	amount := int64(500)
	walletSender := domain.Wallet{UserID: userID, Balance: 1000}
	walletReceiver := domain.Wallet{UserID: toUserID, Balance: 1000}

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(walletSender, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), toUsername).Return(walletReceiver, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, f func(ctx context.Context) error) error {
			return f(ctx)
		},
	)
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), gomock.Any()).Return(nil).Times(2)
	mockStorage.EXPECT().CreateTransactions(gomock.Any(), userID, toUserID, amount).Return(nil)

	err := u.SendCoin(ctx, toUsername, userID, amount)
	assert.NoError(t, err)
}

func TestSendCoin_SelfTransferError(t *testing.T) {
	ctrl, mockStorage, _, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	toUsername := "self"
	walletSender := domain.Wallet{UserID: userID, Balance: 1000}

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(walletSender, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), toUsername).Return(walletSender, nil)

	err := u.SendCoin(ctx, toUsername, userID, 500)
	assert.ErrorIs(t, err, errors.ErrSelfTransfer)
}

func TestSendCoin_InsufficientFundsError(t *testing.T) {
	ctrl, mockStorage, _, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	toUserID := uuid.New()
	toUsername := "receiver"
	walletSender := domain.Wallet{UserID: userID, Balance: 100}
	walletReceiver := domain.Wallet{UserID: toUserID, Balance: 1000}

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(walletSender, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), toUsername).Return(walletReceiver, nil)

	err := u.SendCoin(ctx, toUsername, userID, 500)
	assert.ErrorIs(t, err, errors.ErrInsufficientFound)
}

func TestSendCoin_TransactionFailure(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	toUserID := uuid.New()
	toUsername := "receiver"
	amount := int64(500)
	walletSender := domain.Wallet{UserID: userID, Balance: 1000}
	walletReceiver := domain.Wallet{UserID: toUserID, Balance: 1000}

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(walletSender, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), toUsername).Return(walletReceiver, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).Return(std_errors.New("transaction error"))

	err := u.SendCoin(ctx, toUsername, userID, amount)
	assert.Error(t, err)
}

func TestSendCoin_UpdateWallet1Error(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	toUserID := uuid.New()
	toUsername := "receiver"
	amount := int64(500)
	walletSender := domain.Wallet{UserID: userID, Balance: 1000}
	walletReceiver := domain.Wallet{UserID: toUserID, Balance: 1000}

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(walletSender, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), toUsername).Return(walletReceiver, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, f func(ctx context.Context) error) error {
			return f(ctx)
		},
	)
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), domain.Wallet{UserID: userID, Balance: 500}).Return(std_errors.New("update wallet sender error"))

	err := u.SendCoin(ctx, toUsername, userID, amount)
	assert.Error(t, err)
}

func TestSendCoin_UpdateWallet2Error(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	toUserID := uuid.New()
	toUsername := "receiver"
	amount := int64(500)
	walletSender := domain.Wallet{UserID: userID, Balance: 1000}
	walletReceiver := domain.Wallet{UserID: toUserID, Balance: 1000}

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(walletSender, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), toUsername).Return(walletReceiver, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, f func(ctx context.Context) error) error {
			return f(ctx)
		},
	)
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), domain.Wallet{UserID: userID, Balance: 500}).Return(nil)
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), domain.Wallet{UserID: toUserID, Balance: 1500}).Return(std_errors.New("update wallet receiver error"))

	err := u.SendCoin(ctx, toUsername, userID, amount)
	assert.Error(t, err)
}

func TestSendCoin_CreateTransactionsError(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	toUserID := uuid.New()
	toUsername := "receiver"
	amount := int64(500)
	walletSender := domain.Wallet{UserID: userID, Balance: 1000}
	walletReceiver := domain.Wallet{UserID: toUserID, Balance: 1000}

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(walletSender, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), toUsername).Return(walletReceiver, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, f func(ctx context.Context) error) error {
			return f(ctx)
		},
	)
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), domain.Wallet{UserID: userID, Balance: 500}).Return(nil)
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), domain.Wallet{UserID: toUserID, Balance: 1500}).Return(nil)
	mockStorage.EXPECT().CreateTransactions(gomock.Any(), userID, toUserID, amount).Return(std_errors.New("create transaction error"))

	err := u.SendCoin(ctx, toUsername, userID, amount)
	assert.Error(t, err)
}

func TestSendCoin_GetWalletByUserIDError(t *testing.T) {
	ctrl, mockStorage, _, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	someError := std_errors.New("get wallet by userID error")

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(domain.Wallet{}, someError)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), gomock.Any()).AnyTimes()

	err := u.SendCoin(ctx, "receiver", userID, 500)
	assert.ErrorIs(t, err, someError)
}

func TestSendCoin_GetWalletByUsernameError(t *testing.T) {
	ctrl, mockStorage, _, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	someError := std_errors.New("get wallet by username error")

	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(domain.Wallet{UserID: userID, Balance: 1000}, nil)
	mockStorage.EXPECT().GetWalletByUsername(gomock.Any(), "receiver").Return(domain.Wallet{}, someError)

	err := u.SendCoin(ctx, "receiver", userID, 500)
	assert.ErrorIs(t, err, someError)
}
