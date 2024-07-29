package grpcx

import (
	"github.com/blackhorseya/ryze/pkg/contextx"
	"google.golang.org/grpc"
)

// Server represents the grpc server.
type Server struct {
	grpcserver *grpc.Server
}

// NewServer creates a new grpc server.
func NewServer() (*Server, error) {
	var server *grpc.Server

	server = grpc.NewServer()

	// TODO: 2024/7/29|sean|implement grpc server here

	return &Server{
		grpcserver: server,
	}, nil
}

// Start begins the server.
func (s *Server) Start(ctx contextx.Contextx) error {
	// TODO: 2024/7/29|sean|implement me
	panic("implement me")
}

// Stop stops the server.
func (s *Server) Stop(ctx contextx.Contextx) error {
	// TODO: 2024/7/29|sean|implement me
	panic("implement me")
}
