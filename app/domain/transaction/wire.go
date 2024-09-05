package transaction

import (
	"github.com/google/wire"
)

// ProviderSet will create a new transaction service.
var ProviderSet = wire.NewSet(
	NewTransactionService,
)
