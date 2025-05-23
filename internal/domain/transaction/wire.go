package transaction

import (
	"github.com/blackhorseya/ryze/internal/infra/storage/pgx"
	"github.com/google/wire"
)

// ProviderSet will create a new transaction service.
var ProviderSet = wire.NewSet(
	NewTransactionService,
	pgx.NewTransactionRepo,
)
