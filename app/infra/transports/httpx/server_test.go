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

func TestOptions_GetAddr(t *testing.T) {
	type fields struct {
		Host string
		Port int
		Mode string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				Host: tt.fields.Host,
				Port: tt.fields.Port,
				Mode: tt.fields.Mode,
			}
			if got := o.GetAddr(); got != tt.want {
				t.Errorf("GetAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
