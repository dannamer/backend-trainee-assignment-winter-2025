package sendcoin_usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
)

func (u *SendCoinUsecase) SendCoin(ctx context.Context, toUsername string, userID uuid.UUID, amount int64) error {
	g, gctx := errgroup.WithContext(ctx)
	var walletSender, walletReceiver domain.Wallet

	g.Go(func() error {
		var err error
		walletSender, err = u.storage.GetWalletByUserID(gctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get wallet: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		var err error
		walletReceiver, err = u.storage.GetWalletByUsername(gctx, toUsername)
		if err != nil {
			return fmt.Errorf("failed to get wallet: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	if walletReceiver.UserID == walletSender.UserID {
		return errors.ErrSelfTransfer
	}

	if walletSender.Balance < amount {
		return errors.ErrInsufficientFound
	}
	walletSender.Balance -= amount
	walletReceiver.Balance += amount

	if err := u.trManager.Do(ctx, func(ctx context.Context) error {
		if err := u.storage.UpdateWallet(ctx, walletSender); err != nil {
			return fmt.Errorf("failed to update wallet: %w", err)
		}
		if err := u.storage.UpdateWallet(ctx, walletReceiver); err != nil {
			return fmt.Errorf("failed to update wallet: %w", err)
		}
		if err := u.storage.CreateTransactions(ctx, walletSender.UserID, walletReceiver.UserID, amount); err != nil {
			return fmt.Errorf("failed to save transactions: %w", err)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
