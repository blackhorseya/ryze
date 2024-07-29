package scan

import (
	"github.com/blackhorseya/ryze/adapter/block/wirex"
	_ "github.com/blackhorseya/ryze/api/block/scan" // import swagger
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
)

// @title Ryze Block Scan API
// @version 0.1.0
// @description Ryze Block Scan API document.
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
type restful struct {
	injector *wirex.Injector
}

func NewService(injector *wirex.Injector) adapterx.Service {
	return &restful{
		injector: injector,
	}
}

func (i *restful) Start(ctx contextx.Contextx) error {
	// TODO: 2024/7/28|sean|add block scan logic here
	// i.injector.BlockService.ScanBlock(&model.ScanBlockRequest{}, stream)

	return nil
}

func (i *restful) AwaitSignal(ctx contextx.Contextx) error {
	// TODO: 2024/7/29|sean|add block scan await signal logic here
	return nil
}
