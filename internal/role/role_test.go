package role_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/fwfurtado/blockchain-go/internal/role"
)

var _ = Describe("Role", func() {
	sender := role.NewSender("A")
	reciever := role.NewReciever("B")

	Describe("Transfer amount", func() {
		Context("from sender to reciever", func() {
			It("should return a transaction", func() {
				transaction := sender.Transfer(reciever, 10)

				Expect(transaction.Sender()).To(Equal(sender.Id()))
				Expect(transaction.Reciever()).To(Equal(reciever.Id()))
				Expect(transaction.Amount()).To(Equal(decimal.NewFromFloat(10)))
			})
		})
	})
})
