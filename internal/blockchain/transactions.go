package blockchain

import (
	"github.com/fwfurtado/blockchain-go/internal/transaction"
)

func (b *Blockchain) AddTransaction(transaction transaction.Transaction) {
	b.openTransactions = append(b.openTransactions, transaction)
}

func (b Blockchain) removeItsTransactions(txs transaction.Transactions) transaction.Transactions {
	output := make(transaction.Transactions, 0)

	for _, tx := range b.openTransactions {
		if !txs.Has(tx) {
			output = append(output, tx)
		}
	}

	return output
}

func (b Blockchain) TotalTransctions() int {
	return len(b.openTransactions)
}
