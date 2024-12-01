package ton

import (
	"context"
	"sync"
	"testing"

	blockB "github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/internal/app/repo"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
)

func TestBlockAdapterImpl_ScanBlock(t *testing.T) {
	type fields struct {
		client         *tonx.Client
		shardLastSeqno sync.Map
	}
	type args struct {
		c       context.Context
		req     repo.ScanBlockRequest
		blockCh chan<- *blockB.Block
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BlockAdapterImpl{
				client:         tt.fields.client,
				shardLastSeqno: tt.fields.shardLastSeqno,
			}
			if err := i.ScanBlock(tt.args.c, tt.args.req, tt.args.blockCh); (err != nil) != tt.wantErr {
				t.Errorf("ScanBlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
