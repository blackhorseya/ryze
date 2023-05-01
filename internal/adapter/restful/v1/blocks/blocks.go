package blocks

import (
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/gin-gonic/gin"
)

// Handle will handle the routing of blocks.
func Handle(g *gin.RouterGroup, biz bb.IBiz) {
	instance := &impl{
		biz: biz,
	}

	g.GET("", instance.ListBlocks)
}

type impl struct {
	biz bb.IBiz
}
