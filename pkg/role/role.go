package role

import (
	"github.com/fwfurtado/blockchain-go/pkg/block"
	"github.com/fwfurtado/blockchain-go/pkg/transaction"
)

type Sender struct {
	id string
}

func NewSender(id string) Sender {
	return Sender{
		id: id,
	}
}

func NewReciever(id string) Reciever {
	return Reciever{
		id: id,
	}
}

func (s Sender) Id() string {
	return s.id
}

func (s Sender) Transfer(receiver Reciever, amount float64) block.Transaction {
	return transaction.New(
		s.id,
		receiver.id,
		amount,
	)
}

type Reciever struct {
	id string
}

func (r Reciever) Id() string {
	return r.id
}
