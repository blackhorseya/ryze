package biz

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/account/biz"
	"github.com/blackhorseya/ryze/entity/domain/account/model"
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
	// TODO: 2024/8/16|sean|实现获取账户信息
	panic("implement me")
}
