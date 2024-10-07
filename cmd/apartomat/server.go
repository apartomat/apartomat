package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	defaultAddr = ":80"

	graphQLPath           = "/graphql"
	graphQLPlaygroundPath = "/pg"
)

type ServerOption func(server *http.Server)

func WithAddr(addr string) ServerOption {
	return func(s *http.Server) {
		s.Addr = addr
	}
}

type Server struct {
	router     *chi.Mux
	prometheus *prometheus.Registry

	withGraphQLPlayground bool
}

func NewServer() *Server {
	return &Server{router: chi.NewRouter()}
}

func (server *Server) Run(ctx context.Context, opts ...ServerOption) {
	var (
		s = http.Server{
			Addr:         defaultAddr,
			Handler:      server.router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}
	)

	for _, opt := range opts {
		opt(&s)
	}

	var (
		done = make(chan bool)
		quit = make(chan os.Signal, 1)
	)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit

		slog.InfoContext(ctx, "Stopping server...")

		shutdCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := s.Shutdown(shutdCtx); err != nil {
			slog.ErrorContext(ctx, "can't stop server", err)
			os.Exit(1)
		}

		close(done)
	}()

	slog.InfoContext(ctx, fmt.Sprintf("Starting server at %s...", s.Addr))

	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "Can't start server", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	if server.withGraphQLPlayground {
		slog.InfoContext(ctx, fmt.Sprintf("Visit http://%s%s for playground", serverHttpAddr(s.Addr), graphQLPlaygroundPath))
	}

	<-done

	slog.InfoContext(ctx, "Buy")
}

func (server *Server) Use(next func(http.Handler) http.Handler) *Server {
	server.router.Use(next)

	return server
}

func (server *Server) WithGraphQLHandler(h http.Handler) *Server {
	server.router.Handle(graphQLPath, h)

	return server
}

func (server *Server) WithGraphQLPlayground() *Server {
	server.withGraphQLPlayground = true
	server.router.Handle(graphQLPlaygroundPath, playground.Handler("GraphQL playground", graphQLPath))

	return server
}

func (server *Server) WithMetrics(h http.Handler) *Server {
	server.router.Handle("/metrics", h)

	return server
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
