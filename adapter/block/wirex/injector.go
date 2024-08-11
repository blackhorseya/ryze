package wirex

import (
	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/entity/domain/block/biz"
)

// Injector is the injector for wirex
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	// other fields
	BlockService biz.BlockServiceServer
}
