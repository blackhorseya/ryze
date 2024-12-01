package platform

import (
	"github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/internal/infra/configx"
	"github.com/blackhorseya/ryze/internal/infra/otelx"
)

// Injector is the injector for wirex
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK

	txClient biz.TransactionServiceClient
}
