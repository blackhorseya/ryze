package errorx

import (
	"github.com/blackhorseya/ryze/pkg/er"
)

var (
	// ErrListBlocks is the error that list blocks failed
	ErrListBlocks = er.New(500100, "list blocks failed", "list blocks failed")
)
