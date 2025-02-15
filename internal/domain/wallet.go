package domain

import (
	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance int64
}
