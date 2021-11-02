package network

import "github.com/fwfurtado/blockchain-go/internal/block"

type Node struct {
	Address string
	Port    int
	blocks  block.Blocks
}

func (n Node) Size() int {
	return len(n.blocks)
}

func (n Node) Chain() block.Blocks {
	return n.blocks
}
