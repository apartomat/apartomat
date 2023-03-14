package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// queryLatencyHook is a pg.QueryHook that collect querying metrics
// to prometheus histogram vector
type queryLatencyHook struct {
	vec *prometheus.HistogramVec
}

func NewQueryLatencyHook(reg *prometheus.Registry) pg.QueryHook {
	var (
		vec = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "sql_query_duration_milliseconds",
				Help:    "",
				Buckets: []float64{10, 50, 90, 100, 300, 500, 800, 1000, 5000, 10000},
			},
			[]string{"query"},
		)
	)

	reg.MustRegister(vec)

	return &queryLatencyHook{vec}
}

func (h *queryLatencyHook) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (h *queryLatencyHook) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	h.vec.WithLabelValues(QueryContext(ctx)).Observe(float64(time.Since(q.StartTime).Nanoseconds()) / 1e6)

	return nil
}
