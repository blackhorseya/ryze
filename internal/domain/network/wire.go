package network

import (
	"github.com/google/wire"
)

// ProviderSet will create a new network service.
var ProviderSet = wire.NewSet(
	NewNetworkService,
)
