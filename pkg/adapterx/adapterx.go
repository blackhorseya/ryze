package adapterx

import (
	"context"
)

// Server is the interface that wraps the basic Serve method.
type Server interface {
	Start(c context.Context) error
	Shutdown(c context.Context) error
}
