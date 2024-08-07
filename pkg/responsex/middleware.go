package responsex

import (
	"errors"
	"net/http"

	"github.com/blackhorseya/ryze/pkg/errorx"
	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware is used to add error handling middleware.
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last().Err

			var e *errorx.Error
			if errors.As(err, &e) {
				c.JSON(e.StatusCode, Response{
					Code:    e.Code,
					Message: e.Error(),
				})
				c.Abort()
			} else {
				c.JSON(http.StatusInternalServerError, Response{
					Code:    500,
					Message: err.Error(),
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
