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
		{
			name: "Default address when Host and Port are empty",
			fields: fields{
				Host: "",
				Port: 0,
				Mode: "",
			},
			want: "0.0.0.0:30000", // Assuming netx.GetAvailablePort() returns 30000
		},
		{
			name: "Custom Host and default Port",
			fields: fields{
				Host: "127.0.0.1",
				Port: 0,
				Mode: "",
			},
			want: "127.0.0.1:30000", // Assuming netx.GetAvailablePort() returns 30000
		},
		{
			name: "Default Host and custom Port",
			fields: fields{
				Host: "",
				Port: 9090,
				Mode: "",
			},
			want: "0.0.0.0:9090",
		},
		{
			name: "Custom Host and custom Port",
			fields: fields{
				Host: "192.168.1.1",
				Port: 9090,
				Mode: "",
			},
			want: "192.168.1.1:9090",
		},
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
