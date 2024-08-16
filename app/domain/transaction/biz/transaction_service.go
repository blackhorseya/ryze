package biz

import (
	"context"
	"strconv"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	txB "github.com/blackhorseya/ryze/entity/domain/transaction/biz"
	txM "github.com/blackhorseya/ryze/entity/domain/transaction/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type txService struct {
	client *tonx.Client
}

// NewTransactionService is used to create a new transaction service
func NewTransactionService(client *tonx.Client) txB.TransactionServiceServer {
	return &txService{
		client: client,
	}
}

func (i *txService) GetTransaction(c context.Context, req *txB.GetTransactionRequest) (*txM.Transaction, error) {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}

func (i *txService) ListTransactions(
	req *txB.ListTransactionsRequest,
	stream txB.TransactionService_ListTransactionsServer,
) error {
	ctx, err := contextx.FromContext(stream.Context())
	if err != nil {
		return err
	}

	ctx, span := otelx.Span(ctx, "transaction.biz.ListTransactions")
	defer span.End()

	var txList []*txM.Transaction

	api := ton.NewAPIClient(i.client).WithRetry()

	{
		var fetchedIDs []ton.TransactionShortInfo
		var after *ton.TransactionID3
		var more = true

		for more {
			block, err2 := api.LookupBlock(ctx, req.Workchain, req.Shard, req.SeqNo)
			if err2 != nil {
				ctx.Error("lookup block error", zap.Error(err2), zap.Any("req", &req))
				return err2
			}

			fetchedIDs, more, err2 = api.GetBlockTransactionsV2(ctx, block, 100, after)
			if err2 != nil {
				ctx.Error("get block transactions error", zap.Error(err2), zap.Any("block", &block))
				return err2
			}

			if more {
				after = fetchedIDs[len(fetchedIDs)-1].ID3()
			}

			for _, id := range fetchedIDs {
				tx, err3 := api.GetTransaction(
					ctx,
					block,
					address.NewAddress(0, byte(block.Workchain), id.Account),
					id.LT,
				)
				if err3 != nil {
					ctx.Error("get transaction error", zap.Error(err3), zap.Any("id", id))
					return err3
				}
				ctx.Debug("get transaction", zap.Any("tx", &tx), zap.String("tx_string", tx.String()))

				txList = append(txList, txM.NewTransactionFromTon(tx))
			}
		}
	}

	err = stream.SetHeader(metadata.New(map[string]string{
		"total": strconv.Itoa(len(txList)),
	}))
	if err != nil {
		ctx.Error("set header error", zap.Error(err))
		return err
	}

	for _, tx := range txList {
		if err = stream.Send(tx); err != nil {
			ctx.Error("send transaction error", zap.Error(err))
			return err
		}
	}

	return nil
}
