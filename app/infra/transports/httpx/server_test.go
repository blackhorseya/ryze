package httpx

import (
	"testing"
	"time"

	"github.com/blackhorseya/ryze/pkg/contextx"
)

func TestNewServer(t *testing.T) {
	server, err := NewServer(Options{
		Host: "",
		Port: 0,
		Mode: "",
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

	time.Sleep(3 * time.Second)

	err = server.Stop(contextx.Background())
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
}
