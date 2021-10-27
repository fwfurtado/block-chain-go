package blockchain

import (
	"errors"

	"github.com/fwfurtado/blockchain-go/pkg/block"
	"github.com/fwfurtado/blockchain-go/pkg/hashing"
)

var ErrorEmptyBlockchain = errors.New("blockchain is empty")

func (b *Blockchain) Mine() (*block.Block, error) {
	previous, ok := b.LastBlock()

	if !ok {
		return nil, ErrorEmptyBlockchain
	}

	proof := generateProofOfWorkBy(previous.Proof)
	previousHash := hashing.From(*previous)

	block := b.addBlock(proof, previousHash)

	return &block, nil
}
