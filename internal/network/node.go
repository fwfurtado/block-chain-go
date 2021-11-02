package network

type Node struct {
	Address string
	Port    int
}

func (n Node) ConnectTo(network Network) {
	network = append(network, n)
}
