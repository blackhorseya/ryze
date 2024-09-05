package account

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/otelx"
	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/account/biz"
	"github.com/blackhorseya/ryze/entity/domain/account/model"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
)

type accountService struct {
	client *tonx.Client
}

// NewAccountService will create a new account service
func NewAccountService(client *tonx.Client) biz.AccountServiceServer {
	return &accountService{
		client: client,
	}
}

func (i *accountService) GetAccount(c context.Context, req *biz.GetAccountRequest) (*model.Account, error) {
	next, span := otelx.Tracer.Start(c, "account.biz.GetAccount")
	defer span.End()

	ctx := contextx.WithContext(c)

	api := ton.NewAPIClient(i.client).WithRetry()
	master, err := api.CurrentMasterchainInfo(next)
	if err != nil {
		ctx.Error("failed to get masterchain info", zap.Error(err))
		return nil, err
	}

	addr, err := address.ParseAddr(req.Address)
	if err != nil {
		ctx.Error("failed to parse address", zap.Error(err), zap.String("address", req.Address))
		return nil, err
	}

	// we use WaitForBlock to make sure block is ready,
	// it is optional but escapes us from liteserver block not ready errors
	res, err := api.WaitForBlock(master.SeqNo).GetAccount(next, master, addr)
	if err != nil {
		ctx.Error("failed to get account", zap.Error(err), zap.Any("address", addr))
		return nil, err
	}
	ctx.Debug("get account from ton", zap.Any("account", res))

	// TODO: 2024/8/16|sean|实现获取账户信息

	return model.NewAccountFromSource(res), nil
}
