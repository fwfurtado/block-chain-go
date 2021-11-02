package blockchain

import (
	"github.com/fwfurtado/blockchain-go/internal/block"
	"github.com/fwfurtado/blockchain-go/internal/network"
	"github.com/fwfurtado/blockchain-go/internal/transaction"
)

type Blockchain struct {
	chain            block.Blocks
	openTransactions transaction.Transactions
	network          network.Network
}

func New() Blockchain {

	genesis := block.CreateGenesis()

	chain := block.Blocks{genesis}

	transactions := make(transaction.Transactions, 0)
	nodes := make(network.Network, 0)

	return Blockchain{
		chain:            chain,
		openTransactions: transactions,
		network:          nodes,
	}
}
