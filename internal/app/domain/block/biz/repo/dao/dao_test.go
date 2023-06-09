package dao

import (
	"reflect"
	"testing"
	"time"

	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	time1 = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
)

func TestNewBlock(t *testing.T) {
	type args struct {
		block *bm.Block
	}
	tests := []struct {
		name string
		args args
		want *Block
	}{
		{
			name: "new block",
			args: args{block: &bm.Block{
				Number:           1,
				Hash:             "0x1",
				ParentHash:       "0x0",
				Nonce:            "",
				Sha3Uncles:       "",
				LogsBloom:        "",
				TransactionsRoot: "",
				StateRoot:        "",
				ReceiptsRoot:     "",
				Miner:            "",
				Difficulty:       0,
				TotalDifficulty:  0,
				ExtraData:        "",
				Size:             0,
				GasLimit:         0,
				GasUsed:          0,
				Timestamp:        timestamppb.New(time1),
				Transactions:     nil,
				Uncles:           nil,
			}},
			want: &Block{
				Number:     1,
				Hash:       common.HexToHash("0x1").Bytes(),
				ParentHash: common.HexToHash("0x0").Bytes(),
				Timestamp:  timestamppb.New(time1).AsTime().UTC(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlock(tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_ToEntity(t *testing.T) {
	type fields struct {
		Number     uint64
		Hash       []byte
		ParentHash []byte
		Timestamp  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *bm.Block
	}{
		{
			name: "ok",
			fields: fields{
				Number:     1,
				Hash:       common.HexToHash("0x1").Bytes(),
				ParentHash: common.HexToHash("0x1").Bytes(),
				Timestamp:  time1,
			},
			want: &bm.Block{
				Number:           1,
				Hash:             common.HexToHash("0x1").Hex(),
				ParentHash:       common.HexToHash("0x1").Hex(),
				Nonce:            "",
				Sha3Uncles:       "",
				LogsBloom:        "",
				TransactionsRoot: "",
				StateRoot:        "",
				ReceiptsRoot:     "",
				Miner:            "",
				Difficulty:       0,
				TotalDifficulty:  0,
				ExtraData:        "",
				Size:             0,
				GasLimit:         0,
				GasUsed:          0,
				Timestamp:        timestamppb.New(time1),
				Transactions:     nil,
				Uncles:           nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Block{
				Number:     tt.fields.Number,
				Hash:       tt.fields.Hash,
				ParentHash: tt.fields.ParentHash,
				Timestamp:  tt.fields.Timestamp,
			}
			if got := b.ToEntity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
