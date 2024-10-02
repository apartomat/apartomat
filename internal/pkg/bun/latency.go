package bun

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type QueryLatencyObserveFunc func(dur time.Duration, query string)

type QueryLatencyHook struct {
	observe QueryLatencyObserveFunc
}

func NewQueryLatencyHook(observe QueryLatencyObserveFunc) *QueryLatencyHook {
	return &QueryLatencyHook{observe}
}

func (h *QueryLatencyHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryLatencyHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	h.observe(time.Now().Sub(event.StartTime), QueryContext(ctx))
}
