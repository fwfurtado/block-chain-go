package transaction

type byAmountDesc Transactions

func (txs byAmountDesc) Len() int {
	return len(txs)
}

func (txs byAmountDesc) Less(i, j int) bool {
	return txs[i].Amount().GreaterThan(txs[j].Amount())
}

func (txs byAmountDesc) Swap(i, j int) {
	txs[i], txs[j] = txs[j], txs[i]
}
