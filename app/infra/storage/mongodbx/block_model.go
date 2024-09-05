package mongodbx

import (
	"time"

	"github.com/blackhorseya/ryze/entity/domain/block/model"
)

type blockDocument struct {
	Timestamp time.Time    `json:"timestamp" bson:"timestamp"`
	Metadata  *model.Block `json:"metadata" bson:"metadata"`
}

func newBlockDocument(item *model.Block) *blockDocument {
	return &blockDocument{
		Timestamp: item.Timestamp.AsTime(),
		Metadata:  item,
	}
}
