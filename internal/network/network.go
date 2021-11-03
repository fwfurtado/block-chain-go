package network

type Network []Node

func (n *Network) Connect(node Node) {
	*n = append(*n, node)
}

func (n Network) SelectNode(compare func(a, b Node) Node) (*Node, bool) {
	if len(n) == 0 {
		return nil, false
	}

	result := n[0]

	for _, node := range n[1:] {

		result = compare(result, node)
	}

	return &result, true
}
