package api

import (
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/auth"
	"github.com/dannamer/backend-trainee-assignment-winter-2025/internal/api/buyitem"
)

type API struct {
	*auth.AuthHandler
	*buyitem.BuyItemHandler
}