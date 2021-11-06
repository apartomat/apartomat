package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) (*zap.Logger, error) {
	zl := zapcore.DebugLevel

	if err := zl.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zl),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
