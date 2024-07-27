package scan

import (
	"github.com/blackhorseya/ryze/adpater/block/wirex"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/gin-gonic/gin"
)

type restful struct {
	injector *wirex.Injector
}

// NewRestful is used to create a new adapterx.Restful
func NewRestful(injector *wirex.Injector) adapterx.Restful {
	return &restful{
		injector: injector,
	}
}

func (i *restful) Start() error {
	// TODO implement me
	panic("implement me")
}

func (i *restful) AwaitSignal() error {
	// TODO implement me
	panic("implement me")
}

func (i *restful) InitRouting() error {
	// TODO implement me
	panic("implement me")
}

func (i *restful) GetRouter() *gin.Engine {
	// TODO implement me
	panic("implement me")
}
