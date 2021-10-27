package transaction

import (
	"sort"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	sender   string
	reciever string
	amount   decimal.Decimal
}

func New(sender, receiver string, amount float64) Transaction {
	return Transaction{
		sender:   sender,
		reciever: receiver,
		amount:   decimal.NewFromFloat(amount),
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

type Transactions []Transaction

type byAmountDesc Transactions

func (txs byAmountDesc) Len() int {
	return len(txs)
}

func (txs byAmountDesc) Less(i, j int) bool {
	return txs[i].amount.GreaterThan(txs[j].amount)
}

func (txs byAmountDesc) Swap(i, j int) {
	txs[i], txs[j] = txs[j], txs[i]
}

func (transactions Transactions) TakeGreatestAmount(n int) Transactions {
	temp := make(Transactions, len(transactions))

	copy(temp, transactions)

	sort.Sort(byAmountDesc(temp))

	if n >= len(temp) {
		return temp
	}

	output := make(Transactions, n)

	for index, tx := range temp[:n] {
		output[index] = tx
	}

	return Transactions(output)
}
