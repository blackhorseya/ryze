//go:build external

package scan

import (
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

	err = service.AwaitSignal(ctx)
	if err != nil {
		t.Fatalf("AwaitSignal() error = %v", err)
	}
}
