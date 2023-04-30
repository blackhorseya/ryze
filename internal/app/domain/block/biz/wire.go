//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ryze/internal/app/domain/block/biz/repo"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBlockBiz(repo repo.IRepo) bb.IBiz {
	panic(wire.Build(testProviderSet))
}
