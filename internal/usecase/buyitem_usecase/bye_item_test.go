package buyitem_usecase_test

import (
	"context"
	std_errors "errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/buyitem_usecase"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/usecase/buyitem_usecase/mocks"
)

func setupTest(t *testing.T) (*gomock.Controller, *mocks.MockbuyItemStorage, *mocks.MocktrManager, *buyitem_usecase.BuyItemUsecase) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockbuyItemStorage(ctrl)
	mockTrManager := mocks.NewMocktrManager(ctrl)

	u := buyitem_usecase.New(mockStorage, mockTrManager)

	return ctrl, mockStorage, mockTrManager, u
}

func TestBuyItem_Success(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	item := "sword"
	wallet := domain.Wallet{UserID: userID, Balance: 2000}
	merch := domain.Merch{ID: uuid.New(), Item: item, Price: 1000}

	mockStorage.EXPECT().GetMerchByItem(gomock.Any(), item).Return(merch, nil)
	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(wallet, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	})
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), gomock.Any()).Return(nil)
	mockStorage.EXPECT().UpsertInventory(gomock.Any(), item, userID).Return(nil)

	err := u.BuyItem(ctx, userID, item)
	assert.NoError(t, err)
}

func TestBuyItem_InsufficientFunds(t *testing.T) {
	ctrl, mockStorage, _, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	item := "sword"
	wallet := domain.Wallet{UserID: userID, Balance: 500}
	merch := domain.Merch{Item: item, Price: 1000}

	mockStorage.EXPECT().GetMerchByItem(gomock.Any(), item).Return(merch, nil)
	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(wallet, nil)

	err := u.BuyItem(ctx, userID, item)
	assert.ErrorIs(t, err, errors.ErrInsufficientFound)
}

func TestBuyItem_UpdateWalletErrors(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	item := "sword"
	wallet := domain.Wallet{UserID: userID, Balance: 2000}
	merch := domain.Merch{Item: item, Price: 1000}

	mockStorage.EXPECT().GetMerchByItem(gomock.Any(), gomock.Any()).Return(merch, nil)
	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), gomock.Any()).Return(wallet, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	})
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), gomock.Any()).Return(std_errors.New("update failed"))

	err := u.BuyItem(ctx, userID, item)
	assert.Error(t, err)
}

func TestBuyItem_UpsertInventoryErrors(t *testing.T) {
	ctrl, mockStorage, mockTrManager, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	item := "sword"
	wallet := domain.Wallet{UserID: userID, Balance: 2000}
	merch := domain.Merch{Item: item, Price: 1000}

	mockStorage.EXPECT().GetMerchByItem(gomock.Any(), gomock.Any()).Return(merch, nil)
	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), gomock.Any()).Return(wallet, nil)
	mockTrManager.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, fn func(ctx context.Context) error) error {
		return fn(ctx)
	})
	mockStorage.EXPECT().UpdateWallet(gomock.Any(), gomock.Any()).Return(nil)
	mockStorage.EXPECT().UpsertInventory(gomock.Any(), item, userID).Return(std_errors.New("inventory update failed"))

	err := u.BuyItem(ctx, userID, item)
	assert.Error(t, err)
}

func TestBuyItem_GetMerchByItemErrors(t *testing.T) {
	ctrl, mockStorage, _, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	item := "sword"

	mockStorage.EXPECT().GetMerchByItem(gomock.Any(), item).Return(domain.Merch{}, std_errors.New("merch not found"))
	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(domain.Wallet{}, nil)

	err := u.BuyItem(ctx, userID, item)
	assert.Error(t, err)
}

func TestBuyItem_GetWalletByUserIDErrors(t *testing.T) {
	ctrl, mockStorage, _, u := setupTest(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := uuid.New()
	item := "sword"

	mockStorage.EXPECT().GetMerchByItem(gomock.Any(), item).Return(domain.Merch{Item: item, Price: 1000}, nil)
	mockStorage.EXPECT().GetWalletByUserID(gomock.Any(), userID).Return(domain.Wallet{}, std_errors.New("wallet not found"))

	err := u.BuyItem(ctx, userID, item)
	assert.Error(t, err)
}
