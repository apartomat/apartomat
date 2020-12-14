package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ztsu/apartomat/api/graphql"
	"log"
	foundation "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultAddr = ":80"

type Option interface {
	Apply(server *foundation.Server)
}

type addrOpt struct {
	addr string
}

func Addr(addr string) addrOpt {
	return addrOpt{addr: addr}
}

func (opt addrOpt) Apply(s *foundation.Server) {
	s.Addr = opt.addr
}

type server struct {
	useCases *graphql.UseCases
	stores   *graphql.Stores
}

func NewServer(useCases *graphql.UseCases) *server {
	return &server{
		useCases: useCases,
	}
}

func (server *server) Run(opts ...Option) {
	bgCtx := context.Background()

	mux := foundation.NewServeMux()

	mux.Handle("/graphql", graphql.Handler(
		server.useCases.CheckAuthToken,
		server.stores,
		graphql.NewRootResolver(server.useCases),
	))

	mux.Handle("/pg", playground.Handler("GraphQL playground", "/graphql"))

	s := foundation.Server{
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
		if err := s.ListenAndServe(); err != nil && err != foundation.ErrServerClosed {
			log.Fatalf("can't start server: %s", err)
		}
	}()

	log.Printf("Visit %s/pg for playground", s.Addr)

	<-done

	log.Print("Buy")
}
