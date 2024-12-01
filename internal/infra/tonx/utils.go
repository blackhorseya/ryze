package tonx

import (
	"fmt"

	"github.com/xssnick/tonutils-go/ton"
)

// GetShardID is used to get the shard id
func GetShardID(shard *ton.BlockIDExt) string {
	return fmt.Sprintf("%d|%d", shard.Workchain, shard.Shard)
}
