package ton

import (
	"context"
	"sync"

	"github.com/blackhorseya/ryze/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/internal/app/repo"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
)

// BlockAdapterImpl is the implementation for block adapter.
type BlockAdapterImpl struct {
	client *tonx.Client

	shardLastSeqno sync.Map
}

// NewBlockAdapterImpl is used to create a new block adapter implementation.
func NewBlockAdapterImpl(client *tonx.Client) *BlockAdapterImpl {
	return &BlockAdapterImpl{
		client: client,
	}
}

// NewBlockAdapter is used to create a new block adapter.
func NewBlockAdapter(impl *BlockAdapterImpl) repo.BlockAdapter {
	return impl
}

//nolint:gocognit // it's fine
func (i *BlockAdapterImpl) ScanBlock(
	c context.Context,
	req repo.ScanBlockRequest,
	blockCh chan<- *biz.Block,
) error {
	ctx, span := contextx.StartSpan(c, "datasource.ton.BlockAdapterImpl.ScanBlock")
	defer span.End()

	api := ton.NewAPIClient(i.client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(i.client.Config)

	// 獲取主鏈資訊
	master, err := api.GetMasterchainInfo(ctx)
	if err != nil {
		ctx.Error("failed to get master-chain info", zap.Error(err))
		return err
	}
	ctx.Info("master proofs chain successfully verified, all data is now safe and trusted!")

	// 綁定單一伺服器的上下文以保持一致性
	stickyContext := api.Client().StickyContext(ctx)

	// 從主鏈獲取所有的分片資訊
	firstShards, err := api.GetBlockShardsInfo(stickyContext, master)
	if err != nil {
		ctx.Error("failed to get block shards info", zap.Error(err), zap.Any("master", &master))
		return err
	}

	// 初始化分片序號的記錄
	for _, shard := range firstShards {
		i.shardLastSeqno.Store(tonx.GetShardID(shard), shard.SeqNo)
	}

	// 持續監聽所有分片上的新區塊
	for {
		select {
		case <-ctx.Done():
			ctx.Info("scan block canceled")
			return nil
		default:
			// 獲取每個 workchain 和 shard 上的新區塊
			currentShards, err2 := api.GetBlockShardsInfo(ctx, master)
			if err2 != nil {
				ctx.Error("failed to get block shards info", zap.Error(err2), zap.Any("master", &master))
				return err2
			}

			for _, shard := range currentShards {
				// 檢查是否有新的區塊
				value, ok := i.shardLastSeqno.Load(tonx.GetShardID(shard))
				if ok && shard.SeqNo <= value.(uint32) {
					continue
				}

				// 更新分片序號
				i.shardLastSeqno.Store(tonx.GetShardID(shard), shard.SeqNo)

				// 傳送新區塊
				block, err3 := biz.NewBlock(shard.Workchain, shard.Shard, shard.SeqNo)
				if err3 != nil {
					ctx.Error("failed to create new block", zap.Error(err3))
					return err3
				}
				// TODO: 2024/12/1|sean|if you need more information, you can use the following code
				if blockCh != nil {
					blockCh <- block
				}

				ctx.Info("new block found", zap.Any("shard", &shard), zap.Any("block", &block))
			}

			// 更新主鏈區塊以繼續監控新地分片區塊
			nextSeqNo := master.SeqNo + 1
			master, err2 = api.WaitForBlock(nextSeqNo).LookupBlock(ctx, master.Workchain, master.Shard, nextSeqNo)
			if err2 != nil {
				ctx.Error("failed to lookup next block", zap.Uint32("seq_no", nextSeqNo), zap.Error(err2))
				return err2
			}
		}
	}
}
