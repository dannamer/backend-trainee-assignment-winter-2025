package info_usecase_test

import (
	"context"
	std_errors "errors"
	"testing"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/info_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/info_usecase/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*gomock.Controller, *mocks.MockinfoStorage, *info_usecase.InfoUsecase) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockinfoStorage(ctrl)
	u := info_usecase.New(mockStorage)

	return ctrl, mockStorage, u
}

func TestGetInfo_Success(t *testing.T) {
	ctrl, mockStorage, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	balance := int64(1000)
	inventory := []domain.Inventory{{Item: "sword"}}
	sentTransactions := []domain.Transaction{{Username: "test", Amount: 500}}
	receivedTransactions := []domain.Transaction{{Username: "test", Amount: 700}}

	mockStorage.EXPECT().GetWalletBalanceByUserID(gomock.Any(), userID).Return(balance, nil)
	mockStorage.EXPECT().GetInventoryByUserID(gomock.Any(), userID).Return(inventory, nil)
	mockStorage.EXPECT().GetSentTransactionsByUserID(gomock.Any(), userID).Return(sentTransactions, nil)
	mockStorage.EXPECT().GetReceivedTransactionsByUserID(gomock.Any(), userID).Return(receivedTransactions, nil)

	info, err := u.GetInfo(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, balance, info.Coin)
	assert.Equal(t, inventory, info.Inventory)
	assert.Equal(t, sentTransactions, info.CoinHistory.Sent)
	assert.Equal(t, receivedTransactions, info.CoinHistory.Received)
}

func TestGetInfo_GetWalletBalanceError(t *testing.T) {
	ctrl, mockStorage, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()

	mockStorage.EXPECT().GetWalletBalanceByUserID(gomock.Any(), userID).Return(int64(0), std_errors.New("balance not found"))
	mockStorage.EXPECT().GetInventoryByUserID(gomock.Any(), userID).Return(nil, nil).AnyTimes()
	mockStorage.EXPECT().GetSentTransactionsByUserID(gomock.Any(), userID).Return(nil, nil).AnyTimes()
	mockStorage.EXPECT().GetReceivedTransactionsByUserID(gomock.Any(), userID).Return(nil, nil).AnyTimes()
	_, err := u.GetInfo(ctx, userID)
	assert.Error(t, err)
}

func TestGetInfo_GetInventoryError(t *testing.T) {
	ctrl, mockStorage, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	balance := int64(1000)

	mockStorage.EXPECT().GetWalletBalanceByUserID(gomock.Any(), userID).Return(balance, nil)
	mockStorage.EXPECT().GetInventoryByUserID(gomock.Any(), userID).Return(nil, std_errors.New("inventory not found"))
	mockStorage.EXPECT().GetSentTransactionsByUserID(gomock.Any(), userID).Return(nil, nil).AnyTimes()
	mockStorage.EXPECT().GetReceivedTransactionsByUserID(gomock.Any(), userID).Return(nil, nil).AnyTimes()
	_, err := u.GetInfo(ctx, userID)
	assert.Error(t, err)
}

func TestGetInfo_GetSentTransactionsError(t *testing.T) {
	ctrl, mockStorage, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	balance := int64(1000)
	inventory := []domain.Inventory{{Item: "sword"}}

	mockStorage.EXPECT().GetWalletBalanceByUserID(gomock.Any(), userID).Return(balance, nil)
	mockStorage.EXPECT().GetInventoryByUserID(gomock.Any(), userID).Return(inventory, nil)
	mockStorage.EXPECT().GetSentTransactionsByUserID(gomock.Any(), userID).Return(nil, std_errors.New("sent transactions not found"))
	mockStorage.EXPECT().GetReceivedTransactionsByUserID(gomock.Any(), userID).Return(nil, nil).AnyTimes()
	_, err := u.GetInfo(ctx, userID)
	assert.Error(t, err)
}

func TestGetInfo_GetReceivedTransactionsError(t *testing.T) {
	ctrl, mockStorage, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	balance := int64(1000)
	inventory := []domain.Inventory{{Item: "sword"}}
	sentTransactions := []domain.Transaction{{Username: "test", Amount: 500}}

	mockStorage.EXPECT().GetWalletBalanceByUserID(gomock.Any(), userID).Return(balance, nil)
	mockStorage.EXPECT().GetInventoryByUserID(gomock.Any(), userID).Return(inventory, nil)
	mockStorage.EXPECT().GetSentTransactionsByUserID(gomock.Any(), userID).Return(sentTransactions, nil)
	mockStorage.EXPECT().GetReceivedTransactionsByUserID(gomock.Any(), userID).Return(nil, std_errors.New("received transactions not found"))

	_, err := u.GetInfo(ctx, userID)
	assert.Error(t, err)
}
