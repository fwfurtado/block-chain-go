package blockchain

import (
	"errors"

	"github.com/fwfurtado/blockchain-go/pkg/hashing"
)

var ErrorBlockChainEmpty = errors.New("blockchain is empty")

type Blockchain struct {
	Chain Blocks
}

func New() Blockchain {

	chain := make(Blocks, 1)

	genesis := createBlock(1, "0")

	chain[0] = genesis

	return Blockchain{
		Chain: chain,
	}
}

func (b *Blockchain) Mine() (*Block, error) {
	previous, ok := b.lastBlock()

	if !ok {
		return nil, ErrorBlockChainEmpty
	}

	proof := generateProofOfWorkBy(previous.Proof)
	previousHash := hashing.From(*previous)

	block := b.create(proof, previousHash)

	return &block, nil
}

func (b Blockchain) lastBlock() (*Block, bool) {
	size := len(b.Chain)

	if size > 0 {
		return &b.Chain[size-1], true
	}

	return nil, false
}

func (b *Blockchain) create(proof int, previousHash string) Block {
	block := createBlock(proof, previousHash)
	b.Chain = append(b.Chain, block)

	return block
}

func (b Blockchain) IsValid() bool {

	for previousIndex, block := range b.Chain[1:] {

		previous := b.Chain[previousIndex]
		previousHash := hashing.From(previous)
		if block.Previous != previousHash {
			return false
		}

		if !solveThePuzzle(previous.Proof, block.Proof) {
			return false
		}

	}

	return true
}
