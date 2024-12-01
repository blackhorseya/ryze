package daemon

import (
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/internal/infra/configx"
	"github.com/blackhorseya/ryze/internal/infra/otelx"
)

// Injector is a struct that contains the necessary fields to inject the daemon.
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK

	blockClient blockB.BlockServiceClient
	txClient    txB.TransactionServiceClient
}
