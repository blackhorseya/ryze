package daemon

import (
	blockS "github.com/blackhorseya/ryze/internal/service/block"
	"github.com/blackhorseya/ryze/internal/shared/configx"
	"github.com/blackhorseya/ryze/internal/shared/otelx"
)

// Injector is a struct that contains the necessary fields to inject the daemon.
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK

	// BlockSvc 直接使用本地 block service，供 daemon 掃區塊使用
	BlockSvc blockS.Service
}
