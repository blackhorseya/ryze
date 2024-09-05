package wirex

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/app/infra/otelx"
)

// Injector is the injector for wirex
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK
}
