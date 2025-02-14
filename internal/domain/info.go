package domain

type Info struct {
	Coin        int
	Inventory   []Inventory
	CoinHistory CoinHistory
}

type CoinHistory struct {
	Received []Transaction
	Sent     []Transaction
}

type Transaction struct {
	Username string
	Amount   int
}
