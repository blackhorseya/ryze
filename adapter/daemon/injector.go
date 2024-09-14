package daemon

import (
	"github.com/blackhorseya/ryze/internal/infra/configx"
	"github.com/blackhorseya/ryze/internal/infra/otelx"
)

// Injector is a struct that contains the necessary fields to inject the daemon.
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK
}
