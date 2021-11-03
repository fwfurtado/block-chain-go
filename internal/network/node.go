package network

import "github.com/fwfurtado/blockchain-go/internal/block"

type Node struct {
	address string
	port    int
	blocks  block.Blocks
}

func (n Node) Chain() block.Blocks {
	return n.blocks
}
