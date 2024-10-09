package bun

import (
	"context"
	"github.com/uptrace/bun"
	"log/slog"
	"time"
)

type LogQueryHook struct {
	logger *slog.Logger
}

func NewLogQueryHook(logger *slog.Logger) *LogQueryHook {
	return &LogQueryHook{logger: logger}
}

func (h *LogQueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *LogQueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	var (
		dur = time.Now().Sub(event.StartTime)
	)

	h.logger.DebugContext(
		ctx,
		"Database query",
		slog.Duration("dur", dur),
		slog.String("query", string(event.Query)),
		slog.String("ctx", QueryContext(ctx)),
	)
}
