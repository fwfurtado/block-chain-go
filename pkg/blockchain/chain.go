package blockchain

import (
	"github.com/fwfurtado/blockchain-go/pkg/block"
)

type Blockchain struct {
	chain        block.Blocks
	Transactions block.Transactions
}

func New() Blockchain {

	genesis := block.CreateGenesis()

	chain := block.Blocks{genesis}

	transactions := make(block.Transactions, 0)

	return Blockchain{
		chain:        chain,
		Transactions: transactions,
	}
}

func (b Blockchain) Length() int {
	return len(b.chain)
}

func (b Blockchain) LastBlock() (*block.Block, bool) {
	return b.chain.LastBlock()
}

func (b *Blockchain) AddTransaction(transaction block.Transaction) {
	b.Transactions = append(b.Transactions, transaction)
}

func (b *Blockchain) addBlock(proof int, previousHash string) block.Block {
	block := block.New(proof, previousHash)
	b.chain = append(b.chain, block)
	b.Transactions = nil

	return block
}

func (b Blockchain) Chain() <-chan block.Block {
	generator := make(chan block.Block)

	go func() {
		for _, block := range b.chain {
			generator <- block
		}

		close(generator)
	}()

	return generator
}
