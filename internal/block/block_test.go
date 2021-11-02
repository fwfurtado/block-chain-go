package block_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/fwfurtado/blockchain-go/internal/block"
)

type tx struct {
	amount float64
}

func (t tx) Sender() string {
	return "sender"
}
func (t tx) Reciever() string {
	return "reciever"
}

func (t tx) Amount() decimal.Decimal {
	return decimal.NewFromFloat(t.amount)
}

var _ = Describe("Block", func() {
	genesis := block.CreateGenesis()

	Describe("Creating", func() {
		Context("genesis block", func() {
			It("should return proof 1 and previous hash 0", func() {

				Expect(genesis.Proof).Should(Equal(1))
				Expect(genesis.Previous).Should(Equal("0"))
				Expect(len(genesis.Transactions)).Should(Equal(0))
				Expect(genesis.Timestamp.Before(time.Now().UTC())).Should(BeTrue())
			})
		})
		Context("a block", func() {
			It("should return a block with a given proof and previous hash ", func() {
				block := block.New(2345, "34525")

				Expect(block.Proof).Should(Equal(2345))
				Expect(block.Previous).Should(Equal("34525"))
				Expect(len(block.Transactions)).Should(Equal(0))
				Expect(block.Timestamp.Before(time.Now().UTC())).Should(BeTrue())
			})
		})
	})

	Describe("Given a slice of blocks", func() {
		Context("with items", func() {

			It("should be possible to get the last element ", func() {
				blocks := block.Blocks{
					genesis,
				}

				last, ok := blocks.LastBlock()

				Expect(ok).Should(BeTrue())
				Expect(*last).Should(BeEquivalentTo(genesis))
			})
		})

		Context("without items", func() {
			blocks := block.Blocks{}

			It("should return a nil last element with nok to read", func() {
				last, ok := blocks.LastBlock()

				Expect(ok).Should(BeFalse())
				Expect(last).Should(BeNil())
			})
		})
	})

	Describe("Given a block", func() {

		It("should be possible to add transactions to it", func() {
			block := block.New(1234, "1234")

			Expect(len(block.Transactions)).To(Equal(0))

			transaction := tx{
				amount: 34,
			}

			block.AddTx(transaction)

			Expect(len(block.Transactions)).To(Equal(1))
			Expect(block.Transactions[0].Sender()).To(Equal(transaction.Sender()))
			Expect(block.Transactions[0].Reciever()).To(Equal(transaction.Reciever()))
			Expect(block.Transactions[0].Amount()).To(Equal(transaction.Amount()))

		})
	})

	Describe("taking N transactions ordered by amount", func() {
		allTransactions := make(block.Transactions, 4)

		BeforeEach(func() {
			allTransactions[0] = tx{amount: 30}
			allTransactions[1] = tx{amount: 10}
			allTransactions[2] = tx{amount: 20}
			allTransactions[3] = tx{amount: 40}

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
