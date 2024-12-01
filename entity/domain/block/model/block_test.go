package model

import (
	"reflect"
	"testing"
)

func TestNewBlock(t *testing.T) {
	type args struct {
		workchain int32
		shard     int64
		seqno     uint32
	}
	tests := []struct {
		name    string
		args    args
		want    *Block
		wantErr bool
	}{
		{
			name: "Valid block",
			args: args{
				workchain: -1,
				shard:     8000000000000000,
				seqno:     1,
			},
			want: &Block{
				Id:        "-1:8000000000000000:1",
				Workchain: -1,
				Shard:     8000000000000000,
				SeqNo:     1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBlock(tt.args.workchain, tt.args.shard, tt.args.seqno)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}
