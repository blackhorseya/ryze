package errorx

import (
	"net/http"

	"github.com/blackhorseya/ryze/pkg/er"
)

var (
	// ErrListBlocks is the error that list blocks failed
	ErrListBlocks = er.New(http.StatusInternalServerError, 500100, "list blocks failed", "list blocks failed")
)
