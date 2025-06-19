//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package block

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
	"github.com/blackhorseya/ryze/internal/domain/block/repository"
)

// ScanBlockOptions 包含所有可選參數
// TODO: 擴充 options 欄位
// 例如 Limit、FromID 等
//
//	type ScanBlockOptions struct {
//		Limit  int
//		FromID string
//	}
type ScanBlockOptions struct{}

// ScanBlockOption 代表可選參數
// Example: WithScanBlockLimit, WithScanBlockFromID
type ScanBlockOption func(*ScanBlockOptions)

// Service 定義區塊服務介面
type Service interface {
	// ScanBlock 以 options pattern 與 channel 傳送區塊
	ScanBlock(c context.Context, blocks chan<- *model.Block, opts ...ScanBlockOption) error

	// GetBlockByID 取得單一區塊
	GetBlockByID(c context.Context, id string) (*model.Block, error)

	// GetLatestBlocks 取得多個最新區塊
	GetLatestBlocks(c context.Context, limit int) ([]*model.Block, error)
}

// serviceImpl 為 Service 介面的實作
// TODO: 實作 serviceImpl struct 與方法
type serviceImpl struct {
	blockRepo repository.BlockRepository 
}

// NewService 建立新的區塊服務
func NewService(blockRepo repository.BlockRepository) Service {
	// TODO: 回傳 serviceImpl 實例
	return nil
}

func (s *serviceImpl) ScanBlock(c context.Context, blocks chan<- *model.Block, opts ...ScanBlockOption) error {
	// TODO: 實作區塊掃描邏輯
	return nil
}

func (s *serviceImpl) GetBlockByID(c context.Context, id string) (*model.Block, error) {
	// TODO: 實作取得區塊邏輯
	return nil, nil
}

func (s *serviceImpl) GetLatestBlocks(c context.Context, limit int) ([]*model.Block, error) {
	// TODO: 實作取得最新區塊邏輯
	return nil, nil
}
