package blocks

import (
	"net/http"

	"github.com/blackhorseya/ryze/internal/pkg/errorx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/blackhorseya/ryze/pkg/response"
	"github.com/gin-gonic/gin"
)

type listBlocksQuery struct {
	Page uint `form:"page,default=1" binding:"omitempty,min=1"`
	Size uint `form:"size,default=10" binding:"omitempty,min=1,max=100"`
}

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
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	var query listBlocksQuery
	err := c.BindQuery(&query)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret, total, err := i.biz.ListBlocks(ctx, bb.ListBlocksCondition{
		Page: query.Page,
		Size: query.Size,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(listBlocksResponse{
		Total: total,
		List:  ret,
	}))
}
