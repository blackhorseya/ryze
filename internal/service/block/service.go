//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package block

import (
	"context"

	"github.com/blackhorseya/ryze/internal/domain/block/model"
	"github.com/blackhorseya/ryze/internal/domain/block/repository"
	"github.com/blackhorseya/ryze/internal/domain/block/service"
)

// Service 定義區塊服務介面
type Service interface {
	// ScanBlock 以 options pattern 與 channel 傳送區塊
	ScanBlock(c context.Context, blocks chan<- *model.Block, opts ...service.ScanBlockOption) error

	// GetBlockByID 取得單一區塊
	GetBlockByID(c context.Context, id string) (*model.Block, error)

	// GetLatestBlocks 取得多個最新區塊
	GetLatestBlocks(c context.Context, limit int) ([]*model.Block, error)
}

// serviceImpl 為 Service 介面的實作
type serviceImpl struct {
	blockRepo    repository.BlockRepository
	blockScanner service.BlockScanner // 新增：依賴 Scanner interface
}

// NewService 建立新的區塊服務
func NewService(blockRepo repository.BlockRepository, blockScanner service.BlockScanner) Service {
	return &serviceImpl{
		blockRepo:    blockRepo,
		blockScanner: blockScanner,
	}
}

func (s *serviceImpl) ScanBlock(c context.Context, blocks chan<- *model.Block, opts ...service.ScanBlockOption) error {
	return s.blockScanner.ScanBlocks(c, blocks, opts...)
}

func (s *serviceImpl) GetBlockByID(c context.Context, id string) (*model.Block, error) {
	// TODO: 實作取得區塊邏輯
	return nil, nil
}

func (s *serviceImpl) GetLatestBlocks(c context.Context, limit int) ([]*model.Block, error) {
	// TODO: 實作取得最新區塊邏輯
	return nil, nil
}
