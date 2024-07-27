package wirex

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
)

// Injector is the injector for wirex
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	// other fields
}
