package er

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware global handle *gin.Context error middleware
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()

			switch err.Err.(type) {
			case *Error:
				appError := err.Err.(*Error)
				c.AbortWithStatusJSON(appError.Status, appError)
			default:
				appError := New(http.StatusInternalServerError, 500999, err.Err.Error(), err.Err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, appError)
			}
		}()

		c.Next()
	}
}
