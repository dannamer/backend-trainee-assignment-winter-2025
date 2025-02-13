package domain

import "github.com/shopspring/decimal"

type Info struct {
	Coin        decimal.Decimal
	Inventory   []Inventory
	CoinHistory CoinHistory
}

type CoinHistory struct {
	Received []Transaction
	Sent     []Transaction
}

type Transaction struct {
	Username string
	Amount decimal.Decimal
}
