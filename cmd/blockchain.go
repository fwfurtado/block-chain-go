package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fwfurtado/blockchain-go/pkg/blockchain"
	"github.com/fwfurtado/blockchain-go/pkg/hashing"
)

type blockOutput struct {
	Proof        int       `json:"proof"`
	Timestamp    time.Time `json:"timestamp"`
	Hash         string    `json:"hash"`
	PreviousHash string    `json:"previous_hash"`
}

type blockchainOutput struct {
	Chain []blockOutput `json:"chain"`
}

func main() {
	shitCoin := blockchain.New()

	shitCoin.Mine()
	shitCoin.Mine()
	shitCoin.Mine()
	shitCoin.Mine()
	shitCoin.Mine()

	bytes, _ := json.MarshalIndent(mapping(shitCoin), "", "	")

	fmt.Println(string(bytes))
	fmt.Println("Our blockchain is valid? ", shitCoin.IsValid())
}

func mapping(bc blockchain.Blockchain) blockchainOutput {

	newChain := make([]blockOutput, len(bc.Chain))

	for index, block := range bc.Chain {
		newBlock := blockOutput{
			Proof:        block.Proof,
			Timestamp:    block.Timestamp,
			PreviousHash: block.Previous,
			Hash:         hashing.From(block),
		}

		newChain[index] = newBlock
	}

	return blockchainOutput{
		Chain: newChain,
	}
}
