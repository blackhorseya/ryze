package blocks

import (
	"net/http"

	"github.com/blackhorseya/ryze/internal/pkg/errorx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/response"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetBlockByHash godoc
// @Summary Get block by hash
// @Description Get block by hash
// @Tags Blocks
// @Accept json
// @Produce json
// @Param hash path string true "block hash"
// @Response 200 {object} response.Response{data=model.Block}
// @Response 500 {object} er.Error
// @Router /v1/blocks/{hash} [get]
func (i *impl) GetBlockByHash(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	hash := common.HexToHash(c.Param("hash")).Bytes()
	ret, err := i.biz.GetBlockByHash(ctx, hash)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
