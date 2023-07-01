package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestNewViper(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *viper.Viper
		wantErr bool
	}{
		{
			name:    "input not exist path then error",
			args:    args{path: "./not-exist-path"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewViper(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewViper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewViper() got = %v, want %v", got, tt.want)
			}
		})
	}
}
