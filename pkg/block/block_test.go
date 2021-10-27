package block_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/fwfurtado/blockchain-go/pkg/block"
)

type tx struct{}

func (t tx) Sender() string {
	return "sender"
}
func (t tx) Reciever() string {
	return "reciever"
}

func (t tx) Amount() decimal.Decimal {
	return decimal.NewFromInt(0)
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

			transaction := tx{}

			block.AddTx(transaction)

			Expect(len(block.Transactions)).To(Equal(1))
			Expect(block.Transactions[0].Sender()).To(Equal(transaction.Sender()))
			Expect(block.Transactions[0].Reciever()).To(Equal(transaction.Reciever()))
			Expect(block.Transactions[0].Amount()).To(Equal(transaction.Amount()))

		})
	})
})
