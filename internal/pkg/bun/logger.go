package bun

import (
	"context"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"time"
)

type ZapLogQueryHook struct {
	logger *zap.Logger
}

func NewZapLoggerQueryHook(logger *zap.Logger) *ZapLogQueryHook {
	return &ZapLogQueryHook{logger: logger}
}

func (h *ZapLogQueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *ZapLogQueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	var (
		dur = time.Now().Sub(event.StartTime)
	)

	h.logger.Debug(event.Query, zap.Duration("dur", dur))
}
