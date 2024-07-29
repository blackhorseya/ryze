package grpcx

import (
	"net"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
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
	logger := contextx.Background().Logger
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpcctxtags.StreamServerInterceptor(),
			grpcprometheus.StreamServerInterceptor,
			grpczap.StreamServerInterceptor(logger),
			grpcrecovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpcctxtags.UnaryServerInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
			grpczap.UnaryServerInterceptor(logger),
			grpcrecovery.UnaryServerInterceptor(),
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
