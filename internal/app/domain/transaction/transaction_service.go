package transaction

import (
	"context"
	"errors"
	"io"
	"strconv"

	"github.com/blackhorseya/ryze/entity/domain/block/model"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	txM "github.com/blackhorseya/ryze/entity/domain/transaction/model"
	"github.com/blackhorseya/ryze/entity/domain/transaction/repo"
	"github.com/blackhorseya/ryze/internal/shared/tonx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

func (i *txService) ProcessBlockTransactions(stream grpc.BidiStreamingServer[model.Block, txM.Transaction]) error {
	ctx, span := contextx.StartSpan(stream.Context(), "transaction.biz.ProcessBlockTransactions")
	defer span.End()

	for {
		block, err := stream.Recv()
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.Canceled {
			return nil
		}
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			ctx.Error("receive block error", zap.Error(err))
			return err
		}
		ctx.Debug("receive block", zap.Any("block_id", block.Id))

		list, err := i.fetchTransactionsByBlock(ctx, block)
		if err != nil {
			ctx.Error("list transactions by block error", zap.Error(err), zap.Any("block", &block))
			return err
		}
		for tx := range list {
			if err = i.transactions.Create(ctx, tx); err != nil {
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

func (i *txService) ProcessBlockTransactionsNonStream(
	block *model.Block,
	stream grpc.ServerStreamingServer[txM.Transaction],
) error {
	ctx, span := contextx.StartSpan(stream.Context(), "transaction.biz.ProcessBlockTransactionsNonStream")
	defer span.End()

	list, err := i.fetchTransactionsByBlock(ctx, block)
	if err != nil {
		ctx.Error("list transactions by block error", zap.Error(err), zap.Any("block", &block))
		return err
	}

	var txns []*txM.Transaction
	for tx := range list {
		if err = i.transactions.Create(ctx, tx); err != nil {
			ctx.Error("create transaction error", zap.Error(err), zap.Any("tx", &tx))
			return err
		}

		if err = stream.Send(tx); err != nil {
			ctx.Error("send transaction error", zap.Error(err), zap.Any("tx", &tx))
			return err
		}

		txns = append(txns, tx)
	}
	stream.SetTrailer(metadata.New(map[string]string{"total": strconv.Itoa(len(txns))}))

	return nil
}

// fetchTransactionsByBlock is used to fetch transactions by block
//
//nolint:gocognit,unparam // ignore this
func (i *txService) fetchTransactionsByBlock(c context.Context, block *model.Block) (chan *txM.Transaction, error) {
	// TODO: 2024/10/5|sean|refactor me
	txChan := make(chan *txM.Transaction)

	go func() {
		defer close(txChan)

		ctx, span := contextx.StartSpan(c, "transaction.biz.fetchTransactionsByBlock")
		defer span.End()

		api := ton.NewAPIClient(i.client, ton.ProofCheckPolicyFast).WithRetry()
		api.SetTrustedBlockFromConfig(i.client.Config)
		stickyContext := api.Client().StickyContext(ctx)

		var fetchedIDs []ton.TransactionShortInfo
		var after *ton.TransactionID3
		var more = true

		for more {
			blockInfo, err := api.LookupBlock(stickyContext, block.Workchain, block.Shard, block.SeqNo)
			if errors.Is(err, context.Canceled) {
				return
			}
			if err != nil {
				ctx.Error("lookup block error", zap.Error(err), zap.Any("block", block))
				return
			}

			fetchedIDs, more, err = api.GetBlockTransactionsV2(stickyContext, blockInfo, 100, after)
			if errors.Is(err, context.Canceled) {
				return
			}
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
