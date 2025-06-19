package block

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
	"github.com/blackhorseya/ryze/internal/service/block"
	"github.com/go-kit/kit/endpoint"
)

// GetBlockByIDRequest 取得單一區塊的請求
// Example: ID string
// type GetBlockByIDRequest struct { ID string }
type GetBlockByIDRequest struct {
	ID string `json:"id"`
}

// GetBlockByIDResponse 取得單一區塊的回應
type GetBlockByIDResponse struct {
	Block *model.Block `json:"block,omitempty"`
	Err   string       `json:"error,omitempty"`
}

// GetLatestBlocksRequest 取得多個最新區塊的請求
type GetLatestBlocksRequest struct {
	Limit int `json:"limit"`
}

// GetLatestBlocksResponse 取得多個最新區塊的回應
type GetLatestBlocksResponse struct {
	Blocks []*model.Block `json:"blocks,omitempty"`
	Err    string         `json:"error,omitempty"`
}

// MakeGetBlockByIDEndpoint 建立取得單一區塊的 endpoint
func MakeGetBlockByIDEndpoint(s block.Service) endpoint.Endpoint {
	return func(c context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBlockByIDRequest)
		block, err := s.GetBlockByID(c, req.ID)
		if err != nil {
			return GetBlockByIDResponse{Err: err.Error()}, nil
		}
		return GetBlockByIDResponse{Block: block}, nil
	}
}

// MakeGetLatestBlocksEndpoint 建立取得多個最新區塊的 endpoint
func MakeGetLatestBlocksEndpoint(s block.Service) endpoint.Endpoint {
	return func(c context.Context, request interface{}) (interface{}, error) {
		req := request.(GetLatestBlocksRequest)
		blocks, err := s.GetLatestBlocks(c, req.Limit)
		if err != nil {
			return GetLatestBlocksResponse{Err: err.Error()}, nil
		}
		return GetLatestBlocksResponse{Blocks: blocks}, nil
	}
}

// Endpoints 收集所有 block 相關 endpoint
// Example: ScanBlockEndpoint endpoint.Endpoint
// type Endpoints struct { ... }
type Endpoints struct {
	GetBlockByID    endpoint.Endpoint
	GetLatestBlocks endpoint.Endpoint
}

// NewEndpoints 建立 block endpoints
func NewEndpoints(s block.Service) Endpoints {
	return Endpoints{
		GetBlockByID:    MakeGetBlockByIDEndpoint(s),
		GetLatestBlocks: MakeGetLatestBlocksEndpoint(s),
	}
}
