package domain

import "github.com/shopspring/decimal"

type Info struct {
	Coin        decimal.Decimal
	Inventory   []Inventory
	CoinHistory CoinHistory
}

type Inventory struct {
	Item     string
	Quantity int
}

type CoinHistory struct {
	Received []Received
	Sent     []Sent
}

type Received struct {
	FromUser string
	Amount   decimal.Decimal
}

type Sent struct {
	ToUser string
	Amount decimal.Decimal
}
