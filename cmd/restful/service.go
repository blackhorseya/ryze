package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ryze/pkg/adapter"
	"github.com/blackhorseya/ryze/pkg/app"
	"github.com/blackhorseya/ryze/pkg/httpx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type impl struct {
	logger     *zap.Logger
	httpserver httpx.Server
}

// NewService returns a new restful service.
func NewService(logger *zap.Logger, httpserver httpx.Server, restful adapter.Restful) (app.Servicer, error) {
	err := restful.InitRouting()
	if err != nil {
		return nil, errors.Wrap(err, "init routing error")
	}

	instance := &impl{
		logger:     logger,
		httpserver: httpserver,
	}

	return instance, nil
}

func (i *impl) Start() error {
	if i.httpserver != nil {
		err := i.httpserver.Start()
		if err != nil {
			return errors.Wrap(err, "start http server error")
		}
	}

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		i.logger.Info("receive a signal", zap.String("signal", sig.String()))

		if i.httpserver != nil {
			err := i.httpserver.Stop()
			if err != nil {
				i.logger.Warn("stop http server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}

	return nil
}
