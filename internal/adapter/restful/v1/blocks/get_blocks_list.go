package blocks

import (
	"github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/gin-gonic/gin"
)

type listBlocksResponse struct {
	Total uint           `json:"total,omitempty"`
	List  []*model.Block `json:"list,omitempty"`
}

// ListBlocks godoc
// @Summary List blocks
// @Description List blocks
// @Tags Blocks
// @Accept json
// @Produce json
// @Param page query int false "page" default(1) minimum(1)
// @Param size query int false "page size" default(10) minimum(1) maximum(100)
// @Response 200 {object} response.Response{data=listBlocksResponse}
// @Response 500 {object} er.Error
// @Router /v1/blocks [get]
func (i *impl) ListBlocks(c *gin.Context) {
	// todo: 2023/5/2|sean|impl me
}
