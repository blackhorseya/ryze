package grpcx

import (
	"google.golang.org/grpc"
)

// Server represents the grpc server.
type Server struct {
	grpcserver *grpc.Server
}

// NewServer creates a new grpc server.
func NewServer() (*Server, error) {
	var grpcserver *grpc.Server
	// TODO: 2024/7/29|sean|implement grpc server here

	return &Server{
		grpcserver: grpcserver,
	}, nil
}
