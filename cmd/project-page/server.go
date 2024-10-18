package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/apartomat/apartomat/api/project-page/graphql"
	"github.com/apartomat/apartomat/internal/project-page"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/common/log"
	"log/slog"
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

func WithAddr(addr string) ServerOption {
	return func(s *http.Server) {
		s.Addr = addr
	}
}

type Server struct {
	publicSiteService *project_page.Service
}

func NewServer(service *project_page.Service) *Server {
	return &Server{service}
}

func (server *Server) Run(opts ...ServerOption) {
	mux := chi.NewRouter()

	mux.Handle("/graphql", graphql.Handler(graphql.NewRootResolver(server.publicSiteService)))

	mux.Handle("/pg", playground.Handler("GraphQL playground", "/graphql"))

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

		slog.Info("Stopping server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			slog.Error("can't stop server", err)
			os.Exit(1)
		}

		close(done)
	}()

	slog.Info(fmt.Sprintf("Starting server at %s...", s.Addr))

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			slog.Error("Can't start server: %s", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	log.Info(fmt.Sprintf("Visit http://%s/pg for playground", serverHttpAddr(s.Addr)))

	<-done

	log.Info("Buy")
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
