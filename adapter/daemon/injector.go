package daemon

import (
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	configx2 "github.com/blackhorseya/ryze/internal/app/infra/configx"
	"github.com/blackhorseya/ryze/internal/app/infra/otelx"
)

// Injector is a struct that contains the necessary fields to inject the daemon.
type Injector struct {
	C     *configx2.Configuration
	A     *configx2.Application
	OTelx *otelx.SDK

	blockClient blockB.BlockServiceClient
	txClient    txB.TransactionServiceClient
}
