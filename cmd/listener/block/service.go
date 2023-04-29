package main

import (
	"github.com/blackhorseya/ryze/pkg/adapter"
	"go.uber.org/zap"
)

// Service is a application service
type Service struct {
	logger   *zap.Logger
	listener adapter.Listener
}

// NewService serve caller to create service instance
func NewService(logger *zap.Logger) (*Service, error) {
	svc := &Service{
		logger: logger.With(zap.String("type", "service")),
	}

	return svc, nil
}
