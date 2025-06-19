//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package service

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
)

// BlockScanner 定義區塊掃描契約（僅 interface，不實作）
type BlockScanner interface {
	ScanBlocks(c context.Context, blocks chan<- *model.Block, opts ...ScanBlockOption) error
}

// ScanBlockOptions 包含所有可選參數
type ScanBlockOptions struct {
	StartHeight uint32
	EndHeight   uint32
	Workchain   *int32
}

type ScanBlockOption func(*ScanBlockOptions)

// Option functions
func WithScanBlockStartHeight(height uint32) ScanBlockOption {
	return func(opts *ScanBlockOptions) {
		opts.StartHeight = height
	}
}

func WithScanBlockEndHeight(height uint32) ScanBlockOption {
	return func(opts *ScanBlockOptions) {
		opts.EndHeight = height
	}
}

func WithScanBlockWorkchain(workchain int32) ScanBlockOption {
	return func(opts *ScanBlockOptions) {
		opts.Workchain = &workchain
	}
}
