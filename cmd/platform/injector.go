package platform

import (
	"github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	"github.com/blackhorseya/ryze/internal/shared/configx"
	"github.com/blackhorseya/ryze/internal/shared/otelx"
)

// Injector is the injector for wirex
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK

	txClient biz.TransactionServiceClient
}
