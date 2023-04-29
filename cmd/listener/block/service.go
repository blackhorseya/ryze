package main

import (
	"github.com/blackhorseya/ryze/pkg/adapter"
	"github.com/blackhorseya/ryze/pkg/app"
	"go.uber.org/zap"
)

type service struct {
	logger   *zap.Logger
	listener adapter.Listener
}

// NewService serve caller to create service instance
func NewService(logger *zap.Logger) (app.Servicer, error) {
	svc := &service{
		logger: logger.With(zap.String("type", "service")),
	}

	return svc, nil
}

func (s *service) Start() error {
	// TODO implement me
	panic("implement me")
}

func (s *service) AwaitSignal() error {
	// TODO implement me
	panic("implement me")
}
