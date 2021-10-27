package transaction_test

import (
	"github.com/fwfurtado/blockchain-go/pkg/transaction"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe("Transactions", func() {

	allTransactions := make(transaction.Transactions, 4)

	BeforeEach(func() {
		allTransactions[0] = transaction.New("A", "B", 30)
		allTransactions[1] = transaction.New("B", "C", 10)
		allTransactions[2] = transaction.New("A", "C", 20)
		allTransactions[3] = transaction.New("C", "A", 40)

	})

	Describe("taking N transactions ordered by amount", func() {
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
