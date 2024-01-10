package bun

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"
	"time"
)

// QueryLatencyHook is a pg.QueryHook that collect querying metrics
// to prometheus histogram vector
type QueryLatencyHook struct {
	vec *prometheus.HistogramVec
}

func NewQueryLatencyHook(reg *prometheus.Registry) *QueryLatencyHook {
	var (
		vec = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "sql_query_duration_seconds",
				Help:    "",
				Buckets: []float64{0.10, 0.2, 0.25, 0.3, 0.5, 1, 2, 2.5, 3, 5, 10},
			},
			[]string{"query"},
		)
	)

	reg.MustRegister(vec)

	return &QueryLatencyHook{vec}
}

func (h *QueryLatencyHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryLatencyHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	var (
		dur = time.Now().Sub(event.StartTime)
	)

	h.vec.WithLabelValues(QueryContext(ctx)).Observe(dur.Seconds())
}
