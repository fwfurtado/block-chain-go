package blockchain

import "github.com/fwfurtado/blockchain-go/pkg/hashing"

func (b Blockchain) IsValid() bool {

	for previousIndex, block := range b.chain[1:] {

		previous := b.chain[previousIndex]
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
