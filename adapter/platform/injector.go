package platform

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/entity/domain/transaction/biz"
)

// Injector is the injector for wirex
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK

	txClient biz.TransactionServiceClient
}
