package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"time"
)

type QueryLatencyObserveFunc func(dur time.Duration, query string)

type QueryLatencyHook struct {
	observe QueryLatencyObserveFunc
}

func NewQueryLatencyHook(observe QueryLatencyObserveFunc) pg.QueryHook {
	return &QueryLatencyHook{observe}
}

func (h *QueryLatencyHook) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (h *QueryLatencyHook) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	h.observe(time.Since(q.StartTime), QueryContext(ctx))

	return nil
}
