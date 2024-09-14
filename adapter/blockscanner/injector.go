package blockscanner

import (
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/internal/infra/configx"
	"github.com/blackhorseya/ryze/internal/infra/otelx"
	"github.com/blackhorseya/ryze/pkg/eventx"
)

// Injector is used to inject the dependencies.
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	OTel *otelx.SDK
	Bus  *eventx.EventBus

	// Add more dependencies here
	blockClient blockB.BlockServiceClient
	txClient    txB.TransactionServiceClient
}
