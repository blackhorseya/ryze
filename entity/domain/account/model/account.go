package model

import (
	"github.com/xssnick/tonutils-go/tlb"
)

// NewAccountFromSource is used to create a new Account from tlb.Account
func NewAccountFromSource(value *tlb.Account) *Account {
	return &Account{
		Address:  value.State.Address.Data(),
		Balance:  value.State.Balance.String(),
		IsActive: value.IsActive,
	}
}
