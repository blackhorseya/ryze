//go:build external

package scan

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/spf13/viper"
)

func TestRun(t *testing.T) {
	service, clean, err := New(viper.New())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer clean()

	c := context.Background()

	err = service.Start(c)
	if err != nil {
		t.Fatalf("Start() error = %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	<-signalChan

	err = service.Shutdown(c)
	if err != nil {
		t.Fatalf("AwaitSignal() error = %v", err)
	}
}
