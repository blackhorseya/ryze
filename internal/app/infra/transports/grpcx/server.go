package grpcx

import (
	"net"

	"github.com/blackhorseya/ryze/internal/shared/configx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// InitServers define register handler
type InitServers func(s *grpc.Server)

// Server represents the grpc server.
type Server struct {
	grpcserver *grpc.Server
	addr       string
}

// NewServer creates a new grpc server.
func NewServer(app *configx.Application, init InitServers) (*Server, error) {
	logger := zap.L()
	server := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	init(server)

	return &Server{
		grpcserver: server,
		addr:       app.GRPC.GetAddr(),
	}, nil
}

// Start begins the server.
func (s *Server) Start(ctx contextx.Contextx) error {
	go func() {
		ctx.Info("grpc server start", zap.String("addr", s.addr))

		listen, err := net.Listen("tcp", s.addr)
		if err != nil {
			ctx.Fatal("grpc server listen error", zap.Error(err))
		}

		err = s.grpcserver.Serve(listen)
		if err != nil {
			ctx.Fatal("grpc server serve error", zap.Error(err))
		}
	}()

	return nil
}

// Stop stops the server.
func (s *Server) Stop(ctx contextx.Contextx) error {
	s.grpcserver.Stop()

	return nil
}

// GetAddr returns the server address.
func (s *Server) GetAddr() string {
	return s.addr
}
