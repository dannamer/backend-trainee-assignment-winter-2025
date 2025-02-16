package api

import (
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/auth"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/buyitem"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/info"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/sendcoin"
)

type API struct {
	*auth.AuthHandler
	*buyitem.BuyItemHandler
	*sendcoin.SendCoinHandler
	*info.InfoHandler
}
