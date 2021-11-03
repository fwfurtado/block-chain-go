package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	time     time.Time
	sender   string
	reciever string
	amount   decimal.Decimal
}

func New(sender, receiver string, amount float64) Transaction {
	return Transaction{
		sender:   sender,
		reciever: receiver,
		amount:   decimal.NewFromFloat(amount),
		time:     time.Now().UTC(),
	}
}

func (tx Transaction) Sender() string {
	return tx.sender
}

func (tx Transaction) Reciever() string {
	return tx.reciever
}

func (tx Transaction) Amount() decimal.Decimal {
	return tx.amount
}

func (tx Transaction) Equal(other Transaction) bool {
	return tx.Amount().Equal(other.Amount()) && tx.Sender() == other.Sender() && tx.Reciever() == other.Reciever()
}
