package v1

import (
	"github.com/blackhorseya/ryze/internal/adapter/restful/v1/blocks"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/gin-gonic/gin"
)

// Handle will handle the routing of v1.
func Handle(g *gin.RouterGroup, biz bb.IBiz) {
	blocks.Handle(g.Group("blocks"), biz)
}
