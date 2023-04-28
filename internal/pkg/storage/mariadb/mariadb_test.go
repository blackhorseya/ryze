package mariadb

import (
	"reflect"
	"testing"

	"github.com/golang-migrate/migrate/v4"
)

func TestNewMigration(t *testing.T) {
	type args struct {
		o *Options
	}
	tests := []struct {
		name    string
		args    args
		want    *migrate.Migrate
		wantErr bool
	}{
		{
			name: "new migration",
			args: args{o: &Options{
				URL:    "",
				Debug:  false,
				Conns:  0,
				Source: "github://blackhorseya/ryze/scripts/migrations",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMigration(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMigration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMigration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
