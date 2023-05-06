package blocks

import (
	"net/http"

	"github.com/blackhorseya/ryze/pkg/response"
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
	// todo: 2023/5/7|sean|impl me
	c.JSON(http.StatusOK, response.OK)
}
