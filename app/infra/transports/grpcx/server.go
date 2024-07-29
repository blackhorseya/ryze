package grpcx

import (
	"net"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server represents the grpc server.
type Server struct {
	grpcserver *grpc.Server
	addr       string
}

// NewServer creates a new grpc server.
func NewServer(app *configx.Application) (*Server, error) {
	var server *grpc.Server

	server = grpc.NewServer()

	// TODO: 2024/7/29|sean|implement grpc server here

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
	s.grpcserver.GracefulStop()

	return nil
}
