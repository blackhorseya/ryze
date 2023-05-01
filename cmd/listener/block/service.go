package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ryze/pkg/adapter"
	"github.com/blackhorseya/ryze/pkg/app"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type service struct {
	logger   *zap.Logger
	listener adapter.Listener
	migrate  *migrate.Migrate
}

// NewService serve caller to create service instance
func NewService(logger *zap.Logger, listener adapter.Listener, migrate *migrate.Migrate) (app.Servicer, error) {
	svc := &service{
		logger:   logger.With(zap.String("type", "service")),
		listener: listener,
		migrate:  migrate,
	}

	return svc, nil
}

func (s *service) Start() error {
	if s.migrate != nil {
		err := s.migrate.Up()
		if err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {
				return errors.Wrap(err, "migrate up error")
			}
		}
	}

	if s.listener != nil {
		err := s.listener.Start()
		if err != nil {
			return errors.Wrap(err, "listener start error")
		}
	}

	return nil
}

func (s *service) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		s.logger.Info("receive a signal", zap.String("signal", sig.String()))

		if s.listener != nil {
			err := s.listener.Stop()
			if err != nil {
				s.logger.Warn("stop listener error", zap.Error(err))
			}
		}

		os.Exit(0)
	}

	return nil
}
