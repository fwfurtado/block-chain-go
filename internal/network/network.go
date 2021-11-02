package network

type Network []Node

func (n *Network) Connect(node Node) {
	*n = append(*n, node)
}
