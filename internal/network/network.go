package network

type Network []Node

func (n *Network) Connect(node Node) {
	*n = append(*n, node)
}

func (n Network) MaxNode() Node {
	max := n[0]

	for _, node := range n {

		if node.Size() > max.Size() {
			max = node
		}
	}

	return max
}
