package blockchain

import (
	"time"
)

type Block struct {
	Proof     int
	Timestamp time.Time
	Previous  string
}

type Blocks []Block

func createBlock(proof int, previousHash string) Block {
	return Block{
		Proof:     proof,
		Timestamp: time.Now().UTC(),
		Previous:  previousHash,
	}
}
