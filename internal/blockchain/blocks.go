package blockchain

import (
	"github.com/fwfurtado/blockchain-go/internal/block"
	"github.com/fwfurtado/blockchain-go/internal/transaction"
)

func (b *Blockchain) addBlock(proof int, previousHash string, txs transaction.Transactions) block.Block {
	newBlock := block.New(proof, previousHash)
	b.chain = append(b.chain, newBlock)
	newBlock.ReplaceTxs(txs)

	b.openTransactions = b.removeItsTransactions(txs)

	return newBlock
}

func (b Blockchain) Length() int {
	return len(b.chain)
}

func (b Blockchain) LastBlock() (*block.Block, bool) {
	return b.chain.LastBlock()
}

func (b Blockchain) BlockStream() <-chan block.Block {
	generator := make(chan block.Block)

	go func() {
		for _, block := range b.chain {
			generator <- block
		}

		close(generator)
	}()

	return generator
}
