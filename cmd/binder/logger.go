package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) (*zap.Logger, error) {
	var (
		config = zap.Config{
			Development:      false,
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
		}

		zl = zapcore.DebugLevel
	)

	if err := zl.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	config.Level = zap.NewAtomicLevelAt(zl)

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
