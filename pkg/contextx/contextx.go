package contextx

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.NewExample())
}

// Contextx extends google's context to support logging methods.
type Contextx struct {
	context.Context
	*zap.Logger
}

// Background returns a non-nil, empty Contextx. It is never canceled, has no values, and has no deadline.
// Deprecated: Use WithContext instead.
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
		Logger:  zap.L(),
	}
}

// StartSpan starts a new span with the given name and options.
func StartSpan(c context.Context, spanName string, opts ...trace.SpanStartOption) (Contextx, trace.Span) {
	tracer := otel.Tracer("contextx")
	c, span := tracer.Start(c, spanName, opts...)
	return WithContext(c), span
}
