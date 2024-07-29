package grpcx

import (
	"testing"
	"time"

	"github.com/blackhorseya/ryze/app/infra/configx"
	"github.com/blackhorseya/ryze/pkg/contextx"
)

func TestNewServer(t *testing.T) {
	server, err := NewServer(&configx.Application{})
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
