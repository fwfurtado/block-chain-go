package blockchain

import "github.com/fwfurtado/blockchain-go/internal/network"

func (b Blockchain) AddNode(node network.Node) {
	b.network.Connect(node)
}
