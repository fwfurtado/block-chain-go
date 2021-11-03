package blockchain_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fwfurtado/blockchain-go/internal/blockchain"
	"github.com/fwfurtado/blockchain-go/internal/transaction"
)

func newTx(amount float64) transaction.Transaction {
	return transaction.New(
		uuid.NewString(),
		uuid.NewString(),
		amount,
	)
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
				Expect(chain.TotalTransctions()).Should(Equal(0))
			})

			It("should be valid chain", func() {
				Expect(chain.IsValid()).Should(BeTrue())
			})
		})
	})

})
