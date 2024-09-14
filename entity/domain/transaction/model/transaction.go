package model

import (
	"github.com/xssnick/tonutils-go/tlb"
)

// NewTransactionFromTon is used to create a new transaction from ton
func NewTransactionFromTon(value *tlb.Transaction) *Transaction {
	tx := &Transaction{Id: value.Hash}

	if value.IO.In != nil {
		tx.From = value.IO.In.Msg.SenderAddr().Data()

		if value.IO.In.MsgType == tlb.MsgTypeInternal {
			tx.Amount = value.IO.In.AsInternal().Amount.Nano().Int64()
		}
	}

	// TODO: 2024/8/12|sean|fill more fields
	return tx
}
