//go:generate wire
//go:build wireinject

package biz

import (
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBlockBiz() (bb.IBiz, error) {
	panic(wire.Build(testProviderSet))
}
