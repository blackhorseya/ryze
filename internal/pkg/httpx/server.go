package httpx

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/ryze/internal/pkg/config"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/httpx"
	"github.com/blackhorseya/ryze/pkg/netx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	gin.SetMode(cfg.HTTP.Mode)

	return gin.New()
}

type server struct {
	host string
	port int

	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

func NewServer(cfg *config.Config, logger *zap.Logger, router *gin.Engine) httpx.Server {
	return &server{
		host:       cfg.HTTP.Host,
		port:       cfg.HTTP.Port,
		logger:     logger,
		router:     router,
		httpServer: http.Server{},
	}
}

func (s *server) Start() error {
	if s.host == "" {
		s.host = "0.0.0.0"
	}

	if s.port == 0 {
		s.port = netx.GetAvailablePort()
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	s.logger.Info("http server starting...")

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatal("start http server error", zap.Error(err))
			return
		}
	}()

	s.logger.Info("http server started", zap.String("addr", addr))

	return nil
}

func (s *server) Stop() error {
	s.logger.Info("http server stopping...")

	timeout, cancelFunc := contextx.WithTimeout(contextx.Background(), 5*time.Second)
	defer cancelFunc()

	err := s.httpServer.Shutdown(timeout)
	if err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	s.logger.Info("http server stopped")

	return nil
}

// ServerSet declare http server set
var ServerSet = wire.NewSet(NewRouter, NewServer)
