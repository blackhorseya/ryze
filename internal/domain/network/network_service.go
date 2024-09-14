package network

import (
	"context"

	"github.com/blackhorseya/ryze/entity/domain/network/biz"
	"github.com/blackhorseya/ryze/entity/domain/network/model"
	"github.com/blackhorseya/ryze/internal/infra/otelx"
	"github.com/blackhorseya/ryze/internal/infra/tonx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/zap"
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

func (i *networkService) GetNetworkStats(c context.Context, empty *emptypb.Empty) (*model.NetworkStats, error) {
	next, span := otelx.Tracer.Start(c, "network.biz.GetNetworkStats")
	defer span.End()

	ctx := contextx.WithContext(c)

	api := ton.NewAPIClient(i.client).WithRetry()

	stats, err := api.CurrentMasterchainInfo(next)
	if err != nil {
		ctx.Error("failed to get current masterchain info", zap.Error(err))
		return nil, err
	}

	return &model.NetworkStats{
		TotalBlocks:       0,
		TotalTransactions: 0,
		TotalAccounts:     0,
		LatestBlockHeight: stats.SeqNo,
		LatestBlockTime:   nil,
	}, nil
}

func (i *networkService) GetNodeStatus(c context.Context, req *biz.GetNodeStatusRequest) (*model.NodeStatus, error) {
	// TODO: 2024/8/12|sean|implement me
	panic("implement me")
}
