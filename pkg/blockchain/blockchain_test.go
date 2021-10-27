package blockchain_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/fwfurtado/blockchain-go/pkg/block"
	"github.com/fwfurtado/blockchain-go/pkg/blockchain"
	"github.com/fwfurtado/blockchain-go/pkg/hashing"
)

type tx struct {
	sender   string
	reciever string
	amount   float64
}

func (t tx) Sender() string {
	return t.sender
}
func (t tx) Reciever() string {
	return t.reciever
}

func (t tx) Amount() decimal.Decimal {
	return decimal.NewFromFloat(t.amount)
}

func newTx(amount float64) tx {
	return tx{
		amount:   amount,
		sender:   uuid.NewString(),
		reciever: uuid.NewString(),
	}
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

	Describe("Mining a blockchain ", func() {

		var (
			chain   blockchain.Blockchain
			genesis *block.Block
		)

		BeforeEach(func() {
			chain = blockchain.New()
			genesis, _ = chain.LastBlock()
		})

		Context("without transactions", func() {
			It("should add new block on the chain", func() {

				By("start with genesis")

				Expect(chain.Length()).Should(Equal(1))

				By("mine a new block")
				chain.Mine()

				Expect(chain.Length()).Should(Equal(2))

				lastBlock, _ := chain.LastBlock()

				By("should the last block is different to the genesis block")
				Expect(lastBlock.Proof).ShouldNot(Equal(genesis.Proof))
				Expect(lastBlock.Previous).ShouldNot(Equal(genesis.Previous))

				By("must the last block has the previous hash the same of the genesis hash")
				Expect(lastBlock.Previous).Should(Equal(hashing.From(genesis)))
			})
		})
		Context("with less transactions than maximum peer block", func() {

			var (
				transactions block.Transactions
			)

			BeforeEach(func() {
				transactions = block.Transactions{
					newTx(50),
					newTx(30),
					newTx(20),
					newTx(60),
				}
			})

			It("should add a new block and remove mined transactions", func() {
				By("starting blockchain without transactions")
				Expect(len(chain.Transactions)).Should(Equal(0))

				By("add less transactions then maximum per block")
				for _, tx := range transactions {
					chain.AddTransaction(tx)
				}

				By("blockchain transactions should has the same quantity of transaction added")
				Expect(len(chain.Transactions)).Should(Equal(len(transactions)))

				By("should clear the transactions when mine a blockchain with less then minimum transactions")
				mined, _ := chain.Mine()
				Expect(len(chain.Transactions)).Should(Equal(0))

				By("Mined block should have all transactions")
				Expect(len(mined.Transactions)).Should(Equal(len(transactions)))
				for _, tx := range mined.Transactions {
					Expect(transactions.Has(tx)).Should(BeTrue())
				}

			})
		})
	})
})
