package block

import (
	"time"

	"github.com/fwfurtado/blockchain-go/internal/transaction"
)

type Block struct {
	Proof        int
	Timestamp    time.Time
	Previous     string
	transactions transaction.Transactions
}

type Blocks []Block

func New(proof int, previousHash string) Block {
	return Block{
		Proof:        proof,
		Timestamp:    time.Now().UTC(),
		Previous:     previousHash,
		transactions: make(transaction.Transactions, 0),
	}
}

func CreateGenesis() Block {
	return New(1, "0")
}

func (b *Block) AddTx(transaction transaction.Transaction) {
	b.transactions = append(b.transactions, transaction)
}

func (bs Blocks) LastBlock() (*Block, bool) {
	size := len(bs)

	if size > 0 {
		return &bs[size-1], true
	}

	return nil, false
}
