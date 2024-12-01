package httpx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/netx"
	"github.com/blackhorseya/ryze/pkg/responsex"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

// Server is an HTTP server.
type Server struct {
	httpserver *http.Server

	// Router is the gin engine.
	Router *gin.Engine
}

// Options is the server options.
type Options struct {
	// Host is the server host.
	Host string `json:"host" yaml:"host"`

	// Port is the server port.
	Port int `json:"port" yaml:"port"`

	// Mode is the server mode. Default is "release". Options are "debug" and "test".
	Mode string `json:"mode" yaml:"mode"`
}

// GetAddr is used to get the http address.
func (o *Options) GetAddr() string {
	if o.Host == "" {
		o.Host = "0.0.0.0"
	}

	if o.Port == 0 {
		o.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

// NewServer is used to create a new HTTP server.
func NewServer(options Options) (*Server, error) {
	ctx := contextx.Background()

	gin.SetMode(options.Mode)
	router := gin.New()
	router.Use(AddCorsMiddleware())
	router.Use(ginzap.Ginzap(ctx.Logger, time.RFC3339, true))
	router.Use(otelgin.Middleware("http-server"))
	router.Use(responsex.AddErrorHandlingMiddleware())
	router.Use(ginzap.CustomRecoveryWithZap(ctx.Logger, true, func(c *gin.Context, err any) {
		responsex.Err(c, fmt.Errorf("%v", err))
		c.Abort()
	}))

	httpserver := &http.Server{
		Addr:              options.GetAddr(),
		Handler:           router,
		ReadHeaderTimeout: time.Second,
	}

	return &Server{
		httpserver: httpserver,
		Router:     router,
	}, nil
}

// Start begins the server.
func (s *Server) Start(ctx contextx.Contextx) error {
	go func() {
		ctx.Info("http server start", zap.String("addr", s.httpserver.Addr))

		if err := s.httpserver.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Fatal("start http server error", zap.Error(err))
		}
	}()
	return nil
}

// Stop halts the server.
func (s *Server) Stop(ctx contextx.Contextx) error {
	timeout, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	if err := s.httpserver.Shutdown(timeout); err != nil {
		return err
	}
	return nil
}
