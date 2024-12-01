package account

import (
	"github.com/google/wire"
)

// ProviderSet will create a new account service.
var ProviderSet = wire.NewSet(
	NewAccountService,
)
