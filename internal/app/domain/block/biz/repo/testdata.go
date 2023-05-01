package repo

import (
	"time"

	bm "github.com/blackhorseya/ryze/pkg/entity/domain/block/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	block1 = &bm.Block{
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
		Timestamp:        timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
		Transactions:     nil,
		Uncles:           nil,
	}
)
