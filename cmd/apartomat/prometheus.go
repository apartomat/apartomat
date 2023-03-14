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

func PrometheusLatencyMiddleware(reg *prometheus.Registry) func(next http.Handler) http.Handler {
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
