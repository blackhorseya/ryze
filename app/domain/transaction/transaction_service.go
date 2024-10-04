package transaction

import (
	"context"
	"errors"
	"io"
	"strconv"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/block/model"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	txM "github.com/blackhorseya/ryze/entity/domain/transaction/model"
	"github.com/blackhorseya/ryze/entity/domain/transaction/repo"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type txService struct {
	client *tonx.Client

	transactions repo.ITransactionRepo
}

// NewTransactionService is used to create a new transaction service onchain
func NewTransactionService(client *tonx.Client, transactions repo.ITransactionRepo) txB.TransactionServiceServer {
	return &txService{
		client:       client,
		transactions: transactions,
	}
}

func (i *txService) ProcessBlockTransactions(
	stream grpc.BidiStreamingServer[model.Block, txM.Transaction]) error {
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "transaction.biz.ProcessBlockTransactions")
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
		ctx.Debug("receive block", zap.Any("block_id", block.Id))

		list, err := i.FetchTransactionsByBlock(next, block)
		if err != nil {
			ctx.Error("list transactions by block error", zap.Error(err), zap.Any("block", &block))
			return err
		}
		for tx := range list {
			if err = i.transactions.Create(next, tx); err != nil {
				ctx.Error("create transaction error", zap.Error(err), zap.Any("tx", &tx))
				return err
			}

			if err = stream.Send(tx); err != nil {
				ctx.Error("send transaction error", zap.Error(err), zap.Any("tx", &tx))
				return err
			}
		}
	}
}

// FetchTransactionsByBlock is used to fetch transactions by block
func (i *txService) FetchTransactionsByBlock(
	c context.Context,
	block *model.Block,
) (chan *txM.Transaction, error) {
	txChan := make(chan *txM.Transaction)

	go func() {
		defer close(txChan)

		next, span := otelx.Tracer.Start(c, "transaction.biz.FetchTransactionsByBlock")
		defer span.End()

		ctx := contextx.WithContext(c)

		api := ton.NewAPIClient(i.client, ton.ProofCheckPolicyFast).WithRetry()
		api.SetTrustedBlockFromConfig(i.client.Config)
		stickyContext := api.Client().StickyContext(next)

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
				ctx.Debug("get transaction", zap.String("tx_string", tx.String()))

				got := txM.NewTransactionFromTon(tx)
				got.BlockId = block.Id
				got.Timestamp = block.Timestamp
				txChan <- got
			}
		}
	}()

	return txChan, nil
}

func (i *txService) ListTransactions(
	req *txB.ListTransactionRequest,
	stream grpc.ServerStreamingServer[txM.Transaction],
) error {
	ctx, span := contextx.StartSpan(stream.Context(), "transaction.biz.ListTransactions")
	defer span.End()

	cond := repo.ListTransactionsCondition{
		Limit:  int(req.PageSize),
		Offset: int((req.Page - 1) * req.PageSize),
	}
	items, total, err := i.transactions.List(ctx, cond)
	if err != nil {
		ctx.Error("list transactions error", zap.Error(err), zap.Any("cond", &cond))
		return err
	}

	for _, item := range items {
		if err = stream.Send(item); err != nil {
			ctx.Error("send transaction error", zap.Error(err), zap.Any("item", &item))
			return err
		}
	}
	stream.SetTrailer(metadata.New(map[string]string{"total": strconv.Itoa(total)}))

	return nil
}

func (i *txService) ListTransactionsByAccount(
	req *txB.ListTransactionsByAccountRequest,
	stream grpc.ServerStreamingServer[txM.Transaction],
) error {
	ctx, span := contextx.StartSpan(stream.Context(), "transaction.biz.ListTransactionsByAccount")
	defer span.End()

	cond := repo.ListTransactionsCondition{
		Limit:  int(req.PageSize),
		Offset: int((req.Page - 1) * req.PageSize),
	}
	items, total, err := i.transactions.ListByAccount(ctx, req.AccountId, cond)
	if err != nil {
		ctx.Error("list transactions by account error", zap.Error(err))
		return err
	}

	for _, item := range items {
		if err = stream.Send(item); err != nil {
			ctx.Error("send transaction error", zap.Error(err), zap.Any("item", &item))
			return err
		}
	}
	stream.SetTrailer(metadata.New(map[string]string{"total": strconv.Itoa(total)}))

	return nil
}
