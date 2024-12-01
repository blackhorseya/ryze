package usecase

import (
	"context"
)

type (
	Handler interface {
		Handle(c context.Context, input interface{}) (interface{}, error)
	}
)
