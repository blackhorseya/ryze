package adapterx

import (
	"context"

	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/gin-gonic/gin"
)

// Service is the interface that wraps the basic Serve method.
// Deprecated: use Server instead.
type Service interface {
	// Start a service asynchronously.
	Start(ctx contextx.Contextx) error

	// AwaitSignal waits for a signal to shut down the service.
	AwaitSignal(ctx contextx.Contextx) error
}

// Server is the interface that wraps the basic Serve method.
type Server interface {
	Start(c context.Context) error
	Shutdown(c context.Context) error
}

// Restful is the interface that wraps the restful api method.
type Restful interface {
	Service

	// InitRouting init the routing of restful api.
	InitRouting() error

	// GetRouter returns the router of restful api.
	GetRouter() *gin.Engine
}
