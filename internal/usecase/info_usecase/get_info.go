package info_usecase

import (
	"context"
	"fmt"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

func (u *infoUsecase) GetInfo(ctx context.Context, userID uuid.UUID) (domain.Info, error) {
	g, gctx := errgroup.WithContext(ctx)
	var balance decimal.Decimal
	var inventory []domain.Inventory
	var sentTransaction, receivedTransaction []domain.Transaction

	g.Go(func() error {
		var err error
		balance, err = u.storage.GetWalletBalanceByUserID(gctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get balance: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		inventory, err = u.storage.GetInventoryByUserID(gctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get inventory: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		sentTransaction, err = u.storage.GetSentTransactionsByUserID(gctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get sent transaction: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		receivedTransaction, err = u.storage.GetReceivedTransactionsByUserID(gctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get received transaction: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return domain.Info{}, err
	}

	return domain.Info{
		Coin:      balance,
		Inventory: inventory,
		CoinHistory: domain.CoinHistory{
			Received: receivedTransaction,
			Sent:     sentTransaction,
		},
	}, nil
}
