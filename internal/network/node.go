package network

type Node struct {
	Address string
	Port    int
}

type Network []Node

func (n *Network) AddNode(node Node) {
	*n = append(*n, node)
}
