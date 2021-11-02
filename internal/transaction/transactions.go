package transaction

import "sort"

type Transactions []Transaction

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

func (transactions Transactions) TakeGreatestAmount(n int) Transactions {
	temp := make(Transactions, len(transactions))

	copy(temp, transactions)

	sort.Sort(byAmountDesc(temp))

	if n >= len(temp) {
		return temp
	}

	output := make(Transactions, n)

	for index, tx := range temp[:n] {
		output[index] = tx
	}

	return Transactions(output)
}

func (txs Transactions) Has(transaction Transaction) bool {
	for _, tx := range txs {
		if tx.Equal(transaction) {
			return true
		}
	}

	return false
}
