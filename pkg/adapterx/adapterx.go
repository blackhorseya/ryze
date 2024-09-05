package adapterx

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Server is the interface that wraps the basic Serve method.
type Server interface {
	Start(c context.Context) error
	Shutdown(c context.Context) error
}

// Restful is the interface that wraps the restful api method.
type Restful interface {
	Server

	// InitRouting init the routing of restful api.
	InitRouting() error

	// GetRouter returns the router of restful api.
	GetRouter() *gin.Engine
}
