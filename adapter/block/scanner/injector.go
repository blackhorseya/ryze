package scanner

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
)

// Injector is used to inject the dependencies.
type Injector struct {
	C    *configx.Configuration
	A    *configx.Application
	OTel *otelx.SDK

	// Add more dependencies here
	blockClient biz.BlockServiceClient
}
