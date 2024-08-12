package biz

import (
	"context"

	"github.com/blackhorseya/ryze/app/infra/tonx"
	"github.com/blackhorseya/ryze/entity/domain/network/biz"
	"github.com/blackhorseya/ryze/entity/domain/network/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type networkService struct {
	client *tonx.Client
}

// NewNetworkService return a new network service
func NewNetworkService(client *tonx.Client) biz.NetworkServiceServer {
	return &networkService{
		client: client,
	}
}

func (i *networkService) GetNetworkStats(ctx context.Context, empty *emptypb.Empty) (*model.NetworkStats, error) {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}

func (i *networkService) GetNodeStatus(c context.Context, req *biz.GetNodeStatusRequest) (*model.NodeStatus, error) {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}
