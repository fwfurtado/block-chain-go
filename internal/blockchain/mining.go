package blockchain

import (
	"errors"

	"github.com/fwfurtado/blockchain-go/internal/block"
	"github.com/fwfurtado/blockchain-go/internal/hashing"
)

var ErrorEmptyBlockchain = errors.New("blockchain is empty")

func (b *Blockchain) Mine() (*block.Block, error) {
	previous, ok := b.LastBlock()

	if !ok {
		return nil, ErrorEmptyBlockchain
	}

	txs := b.Transactions.TakeGreatestAmount(5)

	proof := generateProofOfWorkBy(previous.Proof)
	previousHash := hashing.From(*previous)

	block := b.addBlock(proof, previousHash, txs)

	return &block, nil
}
