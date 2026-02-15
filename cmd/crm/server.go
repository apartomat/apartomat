package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
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

func (s *Server) Run(ctx context.Context, opts ...ServerOption) {
	var (
		ser = http.Server{
			Addr:              defaultAddr,
			Handler:           s.router,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      60 * time.Second,
			IdleTimeout:       90 * time.Second,
		}
	)

	for _, opt := range opts {
		opt(&ser)
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

		if err := ser.Shutdown(shutdCtx); err != nil {
			slog.ErrorContext(ctx, "can't stop server", slog.Any("err", err))
			os.Exit(1)
		}

		close(done)
	}()

	slog.InfoContext(ctx, fmt.Sprintf("Starting server at %s...", ser.Addr))

	go func() {
		if err := ser.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "Can't start server", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	if s.withGraphQLPlayground {
		slog.InfoContext(ctx, fmt.Sprintf("Visit http://%s%s for playground", serverHttpAddr(ser.Addr), graphQLPlaygroundPath))
	}

	<-done

	slog.InfoContext(ctx, "Buy")
}

func (s *Server) Use(next func(http.Handler) http.Handler) *Server {
	s.router.Use(next)

	return s
}

func (s *Server) Get(pattern string, handlerFn http.HandlerFunc) *Server {
	s.router.Get(pattern, handlerFn)

	return s
}

func (s *Server) WithGraphQLHandler(h http.Handler) *Server {
	s.router.Handle(graphQLPath, h)

	return s
}

func (s *Server) WithGraphQLPlayground() *Server {
	s.withGraphQLPlayground = true
	s.router.Handle(graphQLPlaygroundPath, playground.Handler("GraphQL playground", graphQLPath))

	return s
}

func (s *Server) WithMetrics(h http.Handler) *Server {
	s.router.Handle("/metrics", h)

	return s
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
