package bun

import (
	"context"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"time"
)

type ZapLoggerQueryHook struct {
	logger *zap.Logger
}

func NewZapLoggerQueryHook(logger *zap.Logger) *ZapLoggerQueryHook {
	return &ZapLoggerQueryHook{logger: logger}
}

func (h *ZapLoggerQueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *ZapLoggerQueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	var (
		dur = time.Now().Sub(event.StartTime)
	)

	h.logger.Debug(event.Query, zap.Duration("dur", dur))
}
