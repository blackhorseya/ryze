package grpcx

import (
	"testing"
	"time"

	"github.com/blackhorseya/ryze/internal/shared/configx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"google.golang.org/grpc"
)

func TestNewServer(t *testing.T) {
	server, err := NewServer(&configx.Application{}, func(s *grpc.Server) {
		t.Log("init grpc server")
	})
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	err = server.Start(contextx.Background())
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	time.Sleep(2 * time.Second)

	err = server.Stop(contextx.Background())
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
}
