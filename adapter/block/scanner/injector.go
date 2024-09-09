package scanner

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
)

// Injector is used to inject the dependencies.
type Injector struct {
	C *configx.Configuration
	A *configx.Application
}
