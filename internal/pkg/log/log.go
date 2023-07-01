package log

import (
	"os"

	"github.com/blackhorseya/ryze/internal/pkg/config"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger serve caller to create zap.Logger
func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	err = level.UnmarshalText([]byte(cfg.Log.Level))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	c := zap.NewDevelopmentEncoderConfig()
	c.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(c)
	if cfg.Log.Output == "json" {
		c = zap.NewProductionEncoderConfig()
		c.EncodeTime = zapcore.RFC3339NanoTimeEncoder
		enc = zapcore.NewJSONEncoder(c)
	}

	cores := make([]zapcore.Core, 0, 2)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	zap.ReplaceGlobals(logger)

	logger.Info("logger init success", zap.Any("log", cfg.Log))

	return logger, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewLogger)
