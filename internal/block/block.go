package block

import (
	"time"
)

type Block struct {
	Proof        int
	Timestamp    time.Time
	Previous     string
	Transactions Transactions
}

type Blocks []Block

func New(proof int, previousHash string) Block {
	return Block{
		Proof:        proof,
		Timestamp:    time.Now().UTC(),
		Previous:     previousHash,
		Transactions: make(Transactions, 0),
	}
}

func CreateGenesis() Block {
	return New(1, "0")
}

func (b *Block) AddTx(transaction Transaction) {
	b.Transactions = append(b.Transactions, transaction)
}

func (bs Blocks) LastBlock() (*Block, bool) {
	size := len(bs)

	if size > 0 {
		return &bs[size-1], true
	}

	return nil, false
}
