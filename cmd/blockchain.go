package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fwfurtado/blockchain-go/internal/blockchain"
	"github.com/fwfurtado/blockchain-go/internal/hashing"
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

type blocksOutput []blockOutput

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

	newChain := make(blocksOutput, bc.Length())

	index := 0

	for block := range bc.BlockStream() {
		newBlock := blockOutput{
			Proof:        block.Proof,
			Timestamp:    block.Timestamp,
			PreviousHash: block.Previous,
			Hash:         hashing.From(block),
		}

		newChain[index] = newBlock
		index++
	}

	return blockchainOutput{
		Chain: newChain,
	}
}
