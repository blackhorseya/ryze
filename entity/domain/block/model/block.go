package model

// NewBlock is used to create a new block.
func NewBlock(workchain int32, shard int64, seqno uint32) (*Block, error) {
	return &Block{
		Workchain: workchain,
		Shard:     shard,
		SeqNo:     seqno,
	}, nil
}
