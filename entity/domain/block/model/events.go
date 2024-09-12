package model

// NewBlockEvent is used to create a new block event.
type NewBlockEvent struct {
	BlockID   string
	Workchain int32
	Shard     int64
	SeqNo     uint32
}
