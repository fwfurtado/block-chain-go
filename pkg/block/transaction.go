package block

import "github.com/shopspring/decimal"

type Transaction interface {
	Sender() string
	Reciever() string
	Amount() decimal.Decimal
}

type Transactions []Transaction
