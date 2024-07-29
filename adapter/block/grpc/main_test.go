//go:build external

package grpc

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	service, err := New(viper.New())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	ctx, cancelFunc := contextx.WithCancel(contextx.Background())
	defer cancelFunc()

	err = service.Start(ctx)
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	<-signalChan

	err = service.AwaitSignal(ctx)
	if err != nil {
		t.Fatalf("AwaitSignal() error = %v", err)
	}
}
