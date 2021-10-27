package blockchain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/fwfurtado/blockchain-go/pkg/block"
	"github.com/fwfurtado/blockchain-go/pkg/blockchain"
	"github.com/fwfurtado/blockchain-go/pkg/hashing"
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

var _ = Describe("Blockchain", func() {

	Describe("Creating", func() {

		chain := blockchain.New()

		Context("a new blockchain", func() {
			It("should return a chain with size 1", func() {
				Expect(chain.Length()).Should(Equal(1))
			})

			It("the last block should be the genesis", func() {
				lastBlock, ok := chain.LastBlock()

				Expect(ok).Should(BeTrue())
				Expect(lastBlock.Proof).To((Equal(1)))
				Expect(lastBlock.Previous).To((Equal("0")))
			})

			It("should has an empty list of trasactions", func() {
				Expect(len(chain.Transactions)).Should(Equal(0))
			})

			It("should be valid chain", func() {
				Expect(chain.IsValid()).Should(BeTrue())
			})
		})
	})

	Describe("Mining", func() {

		chain := blockchain.New()
		genesis, _ := chain.LastBlock()
		var lastBlock *block.Block

		Context("a blockchain", func() {
			It("should add new block on the chain", func() {
				Expect(chain.Length()).Should(Equal(1))

				chain.Mine()

				Expect(chain.Length()).Should(Equal(2))

				lastBlock, _ = chain.LastBlock()
			})

			It("should the last block is different to the genesis block", func() {
				Expect(lastBlock.Proof).ShouldNot(Equal(1))
				Expect(lastBlock.Previous).ShouldNot(Equal("0"))
			})

			It("must the last block has the previous hash the same as the hash of genesis", func() {
				Expect(lastBlock.Previous).Should(Equal(hashing.From(genesis)))
			})
		})
	})

	Describe("Mining with transactions", func() {
		chain := blockchain.New()
		transaction := tx{}

		Context("adding transaction", func() {
			It("should start without transactions", func() {
				Expect(len(chain.Transactions)).Should(Equal(0))
			})
			It("should be possible to add transaction to the blockchain", func() {
				chain.AddTransaction(transaction)
				Expect(len(chain.Transactions)).Should(Equal(1))
			})
			It("should clear the transactions when mine", func() {
				chain.Mine()
				Expect(len(chain.Transactions)).Should(Equal(0))
			})
		})
	})
})
