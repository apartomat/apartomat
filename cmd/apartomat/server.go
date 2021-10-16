package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/apartomat/apartomat/api/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultAddr = ":80"

type Option interface {
	Apply(server *http.Server)
}

type addrOpt struct {
	addr string
}

func Addr(addr string) addrOpt {
	return addrOpt{addr: addr}
}

func (opt addrOpt) Apply(s *http.Server) {
	s.Addr = opt.addr
}

type server struct {
	useCases *graphql.UseCases
	loaders  *apartomat.DataLoaders
}

func NewServer(useCases *graphql.UseCases, loaders *apartomat.DataLoaders) *server {
	return &server{
		useCases: useCases,
		loaders:  loaders,
	}
}

func (server *server) Run(opts ...Option) {
	bgCtx := context.Background()
	reg := prometheus.NewRegistry()

	mux := chi.NewRouter()

	mux.Use(PrometheusMiddleware(reg))

	mux.Handle("/graphql", graphql.Handler(
		server.useCases.CheckAuthToken,
		server.loaders,
		graphql.NewRootResolver(server.useCases),
	))

	mux.Handle("/pg", playground.Handler("GraphQL playground", "/graphql"))

	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	s := http.Server{
		Addr:         defaultAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt.Apply(&s)
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

	log.Printf("Visit %s/pg for playground", s.Addr)

	<-done

	log.Print("Buy")
}
