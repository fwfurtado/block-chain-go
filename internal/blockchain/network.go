package blockchain

import (
	"github.com/fwfurtado/blockchain-go/internal/network"
)

func (b Blockchain) AddNode(node network.Node) {
	b.network.Connect(node)
}

func (b Blockchain) Consensus() {
	withBiggestChainLength := func(a, b network.Node) network.Node {
		if a.Chain().Length() > b.Chain().Length() {
			return a
		}

		return b
	}

	biggestNode, ok := b.network.SelectNode(withBiggestChainLength)

	if !ok {
		return
	}

	biggestNodeChain := biggestNode.Chain()

	if biggestNodeChain.Length() > b.Length() {
		b.chain = biggestNodeChain
	}
}
