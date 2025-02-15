package domain

type Info struct {
	Coin        int64
	Inventory   []Inventory
	CoinHistory CoinHistory
}

type CoinHistory struct {
	Received []Transaction
	Sent     []Transaction
}

type Transaction struct {
	Username string
	Amount   int64
}
