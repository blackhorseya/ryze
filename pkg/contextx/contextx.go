package contextx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.NewExample())
}

const (
	// KeyCtx is the key of contextx.
	KeyCtx = "contextx"
)

// Contextx extends google's context to support logging methods.
type Contextx struct {
	context.Context
	*zap.Logger
}

// Background returns a non-nil, empty Contextx. It is never canceled, has no values, and has no deadline.
func Background() Contextx {
	return Contextx{
		Context: context.Background(),
		Logger:  zap.L(),
	}
}

// WithContext returns a copy of parent in which the context is set to ctx.
func WithContext(c context.Context) Contextx {
	return Contextx{
		Context: c,
		Logger:  ctxzap.Extract(c),
	}
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent Contextx, key, val interface{}) Contextx {
	return Contextx{
		Context: context.WithValue(parent.Context, key, val),
		Logger:  parent.Logger,
	}
}

// WithTimeout returns a copy of the parent context with the timeout adjusted to be no later than d.
func WithTimeout(parent Contextx, d time.Duration) (Contextx, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent.Context, d)

	return Contextx{
		Context: ctx,
		Logger:  parent.Logger,
	}, cancel
}

// WithCancel returns a copy of the parent context with a new Done channel.
func WithCancel(parent Contextx) (Contextx, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent.Context)

	return Contextx{
		Context: ctx,
		Logger:  parent.Logger,
	}, cancel
}

// FromGin returns a Contextx from gin.Context.
func FromGin(c *gin.Context) (Contextx, error) {
	value, exists := c.Get(KeyCtx)
	if !exists {
		return Contextx{}, errors.New("contextx not found in gin.Context")
	}

	ctx, ok := value.(Contextx)
	if !ok {
		return Contextx{}, errors.New("contextx type error")
	}

	return ctx, nil
}

// FromContext returns a Contextx from context.Context.
func FromContext(c context.Context) (ctx Contextx, err error) {
	ctx, ok := c.(Contextx)
	if !ok {
		return Contextx{}, fmt.Errorf("invalid context type: %T", c)
	}

	return ctx, nil
}
