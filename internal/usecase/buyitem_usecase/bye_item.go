package buyitem_usecase

import (
	"context"
	"fmt"

	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/domain"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/infrastructure/errors"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (u *buyItemUsecase) BuyItem(ctx context.Context, userID uuid.UUID, item string) error {
	g, gctx := errgroup.WithContext(ctx)
	var merch domain.Merch
	var wallet domain.Wallet

	g.Go(func() error {
		var err error
		merch, err = u.storage.GetMerchByItem(gctx, item)
		if err != nil {
			return fmt.Errorf("failed to get merch: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		var err error
		wallet, err = u.storage.GetWalletByUserID(gctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get wallet: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	if wallet.Balance.LessThan(merch.Price) {
		return errors.ErrInsufficientFound
	}
	wallet.Balance = wallet.Balance.Sub(merch.Price)

	if err := u.trManager.Do(ctx, func(ctx context.Context) error {
		if err := u.storage.UpdateWallet(ctx, wallet); err != nil {
			return fmt.Errorf("failed to update wallet: %w", err)
		}
		if err := u.storage.SaveInventory(ctx, merch.Item, wallet.UserID); err != nil {
			return fmt.Errorf("failed to save inventory: %w", err)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
