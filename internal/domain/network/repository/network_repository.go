//go:generate mockgen -destination=./mock_network_repository.go -package=repository -source=network_repository.go
package repository

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/network/model"
)

// NetworkRepository 定義網路資料存取契約（DDD Repository interface）
type NetworkRepository interface {
	GetStats(c context.Context) (*model.NetworkStats, error)
	GetNodeStatus(c context.Context, nodeID string) (*model.NodeStatus, error)
	GetAllNodes(c context.Context) ([]*model.NodeStatus, error)
	// 其他查詢/儲存方法
}
