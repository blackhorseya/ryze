//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package block

import (
	"context"
	"errors"

	"github.com/blackhorseya/ryze/entity/domain/block/model"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
	"github.com/xssnick/tonutils-go/ton"
)

// ScanBlockOptions 包含所有可選參數
// 可依需求擴充
// Workchain 為 nil 時代表全部 workchain
// EndHeight=0 代表無上限
// StartHeight=0 代表從最新開始
// TODO: 依實際需求擴充
type ScanBlockOptions struct {
	StartHeight uint32
	EndHeight   uint32
	Workchain   *int32
}

// ScanBlockOption 代表可選參數
// Example: WithScanBlockLimit, WithScanBlockFromID
type ScanBlockOption func(*ScanBlockOptions)

// Service 定義區塊服務介面
type Service interface {
	// ScanBlock 以 options pattern 與 channel 傳送區塊
	ScanBlock(c context.Context, blocks chan<- *model.Block, opts ...ScanBlockOption) error

	// GetBlockByID 取得單一區塊
	GetBlockByID(c context.Context, id string) (*model.Block, error)

	// GetLatestBlocks 取得多個最新區塊
	GetLatestBlocks(c context.Context, limit int) ([]*model.Block, error)
}

// serviceImpl 為 Service 介面的實作
type serviceImpl struct {
	tonClient *tonx.Client
}

// NewService 建立新的區塊服務
func NewService(tonClient *tonx.Client) Service {
	return &serviceImpl{
		tonClient: tonClient,
	}
}

func (s *serviceImpl) ScanBlock(c context.Context, blocks chan<- *model.Block, opts ...ScanBlockOption) error {
	var options ScanBlockOptions
	for _, opt := range opts {
		opt(&options)
	}

	api := ton.NewAPIClient(s.tonClient, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(s.tonClient.Config)

	master, err := api.GetMasterchainInfo(c)
	if err != nil {
		return err
	}
	stickyContext := api.Client().StickyContext(c)
	shardLastSeqno := map[string]uint32{}

	firstShards, err := api.GetBlockShardsInfo(stickyContext, master)
	if err != nil {
		return err
	}
	for _, shard := range firstShards {
		shardLastSeqno[tonx.GetShardID(shard)] = shard.SeqNo
	}

	for {
		currentShards, err2 := api.GetBlockShardsInfo(c, master)
		if errors.Is(err2, context.Canceled) {
			return nil
		}
		if err2 != nil {
			return err2
		}
		for _, shard := range currentShards {
			if options.Workchain != nil && shard.Workchain != *options.Workchain {
				continue
			}
			if lastSeqno, ok := shardLastSeqno[tonx.GetShardID(shard)]; ok && shard.SeqNo <= lastSeqno {
				continue
			}
			if options.EndHeight > 0 && shard.SeqNo > options.EndHeight {
				continue
			}
			shardLastSeqno[tonx.GetShardID(shard)] = shard.SeqNo
			block, err3 := model.NewBlock(shard.Workchain, shard.Shard, shard.SeqNo)
			if err3 != nil {
				return err3
			}
			blocks <- block
		}
		nextSeqNo := master.SeqNo + 1
		master, err2 = api.WaitForBlock(nextSeqNo).LookupBlock(c, master.Workchain, master.Shard, nextSeqNo)
		if errors.Is(err2, context.Canceled) {
			return nil
		}
		if err2 != nil {
			return err2
		}
	}
}

func (s *serviceImpl) GetBlockByID(c context.Context, id string) (*model.Block, error) {
	// TODO: 實作取得區塊邏輯
	return nil, nil
}

func (s *serviceImpl) GetLatestBlocks(c context.Context, limit int) ([]*model.Block, error) {
	// TODO: 實作取得最新區塊邏輯
	return nil, nil
}
