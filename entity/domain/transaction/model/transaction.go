package model

import (
	"github.com/xssnick/tonutils-go/tlb"
)

// NewTransactionFromTon is used to create a new transaction from ton
func NewTransactionFromTon(value *tlb.Transaction) *Transaction {
	// TODO: 2024/8/12|sean|fill more fields
	return &Transaction{
		Id:        value.Hash,
		From:      nil,
		To:        nil,
		Amount:    0,
		Timestamp: nil,
	}
}
