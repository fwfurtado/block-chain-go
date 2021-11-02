package blockchain

import (
	"github.com/fwfurtado/blockchain-go/internal/block"
	"github.com/fwfurtado/blockchain-go/internal/transaction"
)

type Blockchain struct {
	chain        block.Blocks
	transactions transaction.Transactions
}

func New() Blockchain {

	genesis := block.CreateGenesis()

	chain := block.Blocks{genesis}

	transactions := make(transaction.Transactions, 0)

	return Blockchain{
		chain:        chain,
		transactions: transactions,
	}
}
