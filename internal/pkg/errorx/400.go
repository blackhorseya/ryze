package errorx

import (
	"github.com/blackhorseya/ryze/pkg/er"
)

var (
	// ErrInvalidPage is the error that page must be greater than 0
	ErrInvalidPage = er.New(400100, "page must be greater than 0", "invalid page")

	// ErrInvalidSize is the error that size must be greater than 0
	ErrInvalidSize = er.New(400101, "size must be greater than 0", "invalid size")
)
