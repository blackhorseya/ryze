package platform

import (
	"github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	configx2 "github.com/blackhorseya/ryze/internal/app/infra/configx"
	"github.com/blackhorseya/ryze/internal/app/infra/otelx"
)

// Injector is the injector for wirex
type Injector struct {
	C     *configx2.Configuration
	A     *configx2.Application
	OTelx *otelx.SDK

	txClient biz.TransactionServiceClient
}
