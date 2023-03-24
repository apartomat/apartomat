package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/apartomat/apartomat/api/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/dataloader"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const defaultAddr = ":80"

type ServerOption func(server *http.Server)

type addrOpt struct {
	addr string
}

func WithAddr(addr string) ServerOption {
	return func(s *http.Server) {
		s.Addr = addr
	}
}

func (opt addrOpt) Apply(s *http.Server) {
	s.Addr = opt.addr
}

type server struct {
	useCases   *apartomat.Apartomat
	loaders    *dataloader.DataLoaders
	prometheus *prometheus.Registry
	logger     *zap.Logger
}

func NewServer(
	useCases *apartomat.Apartomat,
	loaders *dataloader.DataLoaders,
	reg *prometheus.Registry,
	logger *zap.Logger,
) *server {
	return &server{
		useCases:   useCases,
		loaders:    loaders,
		prometheus: reg,
		logger:     logger,
	}
}

func (server *server) Run(opts ...ServerOption) {
	bgCtx := context.Background()

	mux := chi.NewRouter()

	mux.Use(PrometheusLatencyMiddleware(server.prometheus))

	mux.Handle("/graphql", graphql.Handler(
		server.useCases.CheckAuthToken,
		server.loaders,
		graphql.NewRootResolver(server.useCases, server.logger),
		10000,
	))

	mux.Handle("/pg", playground.Handler("GraphQL playground", "/graphql"))

	mux.Handle("/metrics", promhttp.HandlerFor(server.prometheus, promhttp.HandlerOpts{}))

	s := http.Server{
		Addr:         defaultAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(&s)
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit

		log.Print("Stopping server...")

		ctx, cancel := context.WithTimeout(bgCtx, 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Fatalf("can't stop server: %s", err)
		}

		close(done)
	}()

	log.Printf("Starting server at %s...", s.Addr)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("can't start server: %s", err)
		}
	}()

	log.Printf("Visit http://%s/pg for playground", serverHttpAddr(s.Addr))

	<-done

	log.Print("Buy")
}

func serverHttpAddr(addr string) string {
	if !strings.HasPrefix(addr, ":") {
		return addr
	}

	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String() + addr
				}
			}
		}
	}

	return addr
}
