package blockchain

import (
	"github.com/fwfurtado/blockchain-go/internal/network"
)

func (b Blockchain) AddNode(node network.Node) {
	b.network.Connect(node)
}

func (b Blockchain) Consensus() {
	max := b.network.MaxNode()

	if max.Size() > b.Length() {
		b.chain = max.Chain()
	}
}
