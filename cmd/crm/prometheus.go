package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	sqlHistogramVec *prometheus.HistogramVec
)

func init() {
	sqlHistogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "sql_query_duration_seconds",
			Help:    "",
			Buckets: []float64{0.10, 0.2, 0.25, 0.3, 0.5, 1, 2, 2.5, 3, 5, 10},
		},
		[]string{"query"},
	)
}

func observeSql(dur time.Duration, query string) {
	sqlHistogramVec.WithLabelValues(query).Observe(dur.Seconds())
}

func NewMetrics() (prometheus.Registerer, prometheus.Gatherer) {
	reg := prometheus.NewRegistry()

	wr := prometheus.WrapRegistererWith(prometheus.Labels{"service_name": "crm"}, reg)

	wr.MustRegister(sqlHistogramVec)

	return wr, reg
}

func PrometheusLatencyMiddleware(reg prometheus.Registerer) func(next http.Handler) http.Handler {
	var (
		_requestTimer = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "requests_duration_seconds",
			},
			[]string{"method", "path", "status"},
		)

		_requestsTotal = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "requests_total",
				Help: "The total number of requests",
			},
			[]string{"method", "path", "status"},
		)
	)

	reg.MustRegister(_requestTimer, _requestsTotal)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			defer func() {
				ctx := chi.RouteContext(r.Context())
				lvs := []string{
					r.Method,
					strings.TrimRight(ctx.RoutePattern(), "/"),
					strconv.Itoa(rww.Status()),
				}
				_requestsTotal.WithLabelValues(lvs...).Inc()
				_requestTimer.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
			}()

			next.ServeHTTP(rww, r)
		})
	}
}
