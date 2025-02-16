package models

type UserInfo struct {
	Coins     int64
	Inventory []InventoryItem
	Sent      []Transaction
	Received  []Transaction
}

type UserCredentials struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type User struct {
	Id           int64  `db:"id"`
	Login        string `db:"login"`
	PasswordHash string `db:"password_hash"`
	Coins        int64  `db:"coins"`
}

type InventoryItem struct {
	Type     string `db:"type"`
	Quantity int64  `db:"quantity"`
}

type Item struct {
	Id    int64  `db:"id"`
	Type  string `db:"type"`
	Coins int64  `db:"coins"`
}

type Transaction struct {
	Receiver string
	Sender   string
	Amount   int64
}
