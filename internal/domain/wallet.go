package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance decimal.Decimal
}

// func (w *Wallet) AddBalance(amount decimal.Decimal) {
// 	w.Balance = w.Balance.Add(amount)
// }

// func (w *Wallet) SubtractBalance(amount decimal.Decimal) error {
// 	if w.Balance.LessThan(amount) {
// 		return errors.New("insufficient funds")
// 	}
// 	w.Balance = w.Balance.Sub(amount)
// 	return nil
// }
