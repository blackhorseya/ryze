package daemon

import (
	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	blockS "github.com/blackhorseya/ryze/internal/service/block"
	"github.com/blackhorseya/ryze/internal/shared/configx"
	"github.com/blackhorseya/ryze/internal/shared/otelx"
)

// Injector is a struct that contains the necessary fields to inject the daemon.
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK

	blockClient blockB.BlockServiceClient
	txClient    txB.TransactionServiceClient

	// BlockSvc 直接使用本地 block service，供 daemon 掃區塊使用
	BlockSvc blockS.Service
}
