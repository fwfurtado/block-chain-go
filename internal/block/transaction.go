package block

import "github.com/fwfurtado/blockchain-go/internal/transaction"

func (b Block) TotalTransctions() int {
	return len(b.transactions)
}

func (b Block) StreamTransactions() <-chan transaction.Transaction {
	ch := make(chan transaction.Transaction)

	go func() {
		for _, tx := range b.transactions {
			ch <- tx
		}
		close(ch)
	}()

	return ch
}

func (b *Block) ReplaceTxs(transactions transaction.Transactions) {
	b.transactions = transactions
}

func (b *Block) AddTx(transaction transaction.Transaction) {
	b.transactions = append(b.transactions, transaction)
}
