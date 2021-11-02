package transaction_test

import (
	"github.com/fwfurtado/blockchain-go/internal/transaction"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var newTx = func(amount float64) transaction.Transaction {
	return transaction.New(
		"from",
		"to",
		amount,
	)
}

var _ = Describe("Transaction", func() {
	Context("", func() {
		Describe("taking N transactions ordered by amount", func() {
			allTransactions := make(transaction.Transactions, 4)

			BeforeEach(func() {
				allTransactions[0] = newTx(30)
				allTransactions[1] = newTx(10)
				allTransactions[2] = newTx(20)
				allTransactions[3] = newTx(40)

			})
			Context("when total transactions is equal to N", func() {
				It("should return all transactions", func() {
					orderedTransactions := allTransactions.TakeGreatestAmount(4)

					Expect(len(orderedTransactions)).To(Equal(4))
					Expect(orderedTransactions[0].Amount()).To(BeEquivalentTo(decimal.NewFromFloat(40)))
					Expect(orderedTransactions[1].Amount()).To(BeEquivalentTo(decimal.NewFromFloat(30)))
					Expect(orderedTransactions[2].Amount()).To(BeEquivalentTo(decimal.NewFromFloat(20)))
					Expect(orderedTransactions[3].Amount()).To(BeEquivalentTo(decimal.NewFromFloat(10)))

				})
			})

			Context("when total transaction is lower then N", func() {
				It("should return the N transactions with greatest amount ", func() {
					orderedTransactions := allTransactions.TakeGreatestAmount(2)

					Expect(len(orderedTransactions)).To(Equal(2))
					Expect(orderedTransactions[0].Amount()).To(BeEquivalentTo(decimal.NewFromFloat(40)))
					Expect(orderedTransactions[1].Amount()).To(BeEquivalentTo(decimal.NewFromFloat(30)))
				})
			})
		})
	})
})
