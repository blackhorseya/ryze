package contextx

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// AddContextxMiddleware is used to add contextx middleware.
func AddContextxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KeyCtx, WithContext(c.Request.Context()))

		c.Next()
	}
}

// UnaryServerInterceptor is used to create a new unary interceptor
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	// TODO: 2024/8/11|sean|implement me
	panic("implement me")
}

// StreamServerInterceptor is used to create a new stream interceptor
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	// TODO: 2024/8/11|sean|implement me
	panic("implement me")
}
