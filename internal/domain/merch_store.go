package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Merch struct {
	ID    uuid.UUID
	Item  string
	Price decimal.Decimal
}
