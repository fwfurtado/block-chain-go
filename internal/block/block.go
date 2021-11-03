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
