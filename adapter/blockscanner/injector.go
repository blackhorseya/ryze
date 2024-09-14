package blockscanner

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
)

// Injector is used to inject the dependencies.
type Injector struct {
	C    *configx.Configuration
	A    *configx.Application
	OTel *otelx.SDK

	// Add more dependencies here
	blockClient blockB.BlockServiceClient
	txClient    txB.TransactionServiceClient
}
