package errorx

import (
	"net/http"

	"github.com/blackhorseya/ryze/pkg/er"
)

var (
	// ErrContextx is the error that contextx is not found
	ErrContextx = er.New(http.StatusInternalServerError, 500000, "contextx is not found", "contextx is not found")
)

var (
	// ErrListBlocks is the error that list blocks failed
	ErrListBlocks = er.New(http.StatusInternalServerError, 500100, "list blocks failed", "list blocks failed")
)
