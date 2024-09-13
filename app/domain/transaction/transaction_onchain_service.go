package transaction

import (
	"errors"
	"io"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	txM "github.com/blackhorseya/ryze/entity/domain/transaction/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type txServiceOnchain struct {
	client *tonx.Client
}

// NewTransactionService is used to create a new transaction service
func NewTransactionService(client *tonx.Client) txB.TransactionServiceServer {
	return &txServiceOnchain{
		client: client,
	}
}

func (i *txServiceOnchain) ListTransactions(
	req *txB.ListTransactionsRequest,
	stream txB.TransactionService_ListTransactionsServer,
) error {
	c := stream.Context()
	_, span := otelx.Tracer.Start(c, "transaction.biz.ListTransactions")
	defer span.End()

	ctx := contextx.WithContext(c)

	block, err := model.NewBlock(req.Workchain, req.Shard, req.SeqNo)
	if err != nil {
		ctx.Error("new block error", zap.Error(err))
		return err
	}

	list, err := i.ListTransactionsByBlock(ctx, block)
	if err != nil {
		ctx.Error("list transactions by block error", zap.Error(err), zap.Any("block", &block))
		return err
	}
	for tx := range list {
		if err = stream.Send(tx); err != nil {
			ctx.Error("send transaction error", zap.Error(err), zap.Any("tx", &tx))
			return err
		}
	}

	return nil
}

func (i *txServiceOnchain) ProcessBlockTransactions(stream grpc.BidiStreamingServer[model.Block, txM.Transaction]) error {
	c := stream.Context()
	_, span := otelx.Tracer.Start(c, "transaction.biz.ProcessBlockTransactions")
	defer span.End()

	ctx := contextx.WithContext(c)

	for {
		block, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			ctx.Error("receive block error", zap.Error(err))
			return err
		}
		ctx.Debug("receive block", zap.Any("block", &block))

		list, err := i.ListTransactionsByBlock(ctx, block)
		if err != nil {
			ctx.Error("list transactions by block error", zap.Error(err), zap.Any("block", &block))
			return err
		}
		for tx := range list {
			if err = stream.Send(tx); err != nil {
				ctx.Error("send transaction error", zap.Error(err), zap.Any("tx", &tx))
				return err
			}
		}
	}
}

// ListTransactionsByBlock is used to list transactions by block
func (i *txServiceOnchain) ListTransactionsByBlock(
	ctx contextx.Contextx,
	block *model.Block,
) (chan *txM.Transaction, error) {
	txChan := make(chan *txM.Transaction)

	go func() {
		defer close(txChan)

		api := ton.NewAPIClient(i.client, ton.ProofCheckPolicyFast).WithRetry()
		api.SetTrustedBlockFromConfig(i.client.Config)
		stickyContext := api.Client().StickyContext(ctx)

		var fetchedIDs []ton.TransactionShortInfo
		var after *ton.TransactionID3
		var more = true

		for more {
			blockInfo, err := api.LookupBlock(stickyContext, block.Workchain, block.Shard, block.SeqNo)
			if err != nil {
				ctx.Error("lookup block error", zap.Error(err), zap.Any("block", block))
				return
			}

			fetchedIDs, more, err = api.GetBlockTransactionsV2(stickyContext, blockInfo, 100, after)
			if err != nil {
				ctx.Error("get block transactions error", zap.Error(err), zap.Any("blockInfo", blockInfo))
				return
			}

			if more {
				after = fetchedIDs[len(fetchedIDs)-1].ID3()
			}

			for _, id := range fetchedIDs {
				tx, err2 := api.GetTransaction(
					stickyContext,
					blockInfo,
					address.NewAddress(0, byte(blockInfo.Workchain), id.Account),
					id.LT,
				)
				if err2 != nil {
					ctx.Error("get transaction error", zap.Error(err2), zap.Any("id", id))
					return
				}
				ctx.Debug("get transaction", zap.Any("tx", &tx), zap.String("tx_string", tx.String()))

				got := txM.NewTransactionFromTon(tx)
				txChan <- got
			}
		}
	}()

	return txChan, nil
}
