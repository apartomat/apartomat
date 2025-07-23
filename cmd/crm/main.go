package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/apartomat/apartomat/api/crm/graphql"
	"github.com/apartomat/apartomat/api/crm/graphql/dataloaders"
	crmparams "github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/pkg/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		logLevel, _ = logLevel(os.Getenv("LOG_LEVEL"))
	)

	slog.SetDefault(slog.New(
		log.NewAttrHandler(
			slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel, AddSource: false}),
		),
	))

	if len(os.Args) < 2 {
		slog.Error("expect command (run or gen-key-pair)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gen-key-pair":
		genKeyPair()
	case "run":
		run()
	default:
		slog.Info("expect command (run or gen-key-pair)")
		os.Exit(1)
	}
}

func genKeyPair() {
	_, _, err := genPairToFile("apartomat.key")
	if err != nil {
		slog.Error("can't generate key pair", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("done")

	os.Exit(0)
}

func run() {
	//uploader, err := s3.NewS3ImageUploaderWithCred(
	//	ctx,
	//	os.Getenv("S3_ACCESS_KEY_ID"),
	//	os.Getenv("S3_SECRET_ACCESS_KEY"),
	//	os.Getenv("S3_REGION"),
	//	os.Getenv("S3_BUCKET_NAME"),
	//)
	//if err != nil {
	//	log.Fatalf("can't init s3: %s", err)
	//}

	crm, err := InitializeCRM(context.Background())
	if err != nil {
		slog.Error("can't initialize app", slog.Any("err", err))
		os.Exit(1)
	}

	var (
		addr = WithAddr(fmt.Sprintf(":%s", os.Getenv("PORT")))
	)

	if os.Getenv("SERVER_ADDR") != "" {
		addr = WithAddr(os.Getenv("SERVER_ADDR"))
	}

	h := graphql.Handler(
		crm.CheckAuthToken,
		func() *dataloaders.DataLoaders {
			return dataloaders.NewDataLoaders(
				crm.Files,
				crm.Rooms,
				crm.Users,
				crm.Workspaces,
			)
		},
		graphql.NewRootResolver(crm.DB, crm),
		10000,
	)

	reg, gath := NewMetrics()

	NewServer().
		Use(PrometheusLatencyMiddleware(reg)).
		WithGraphQLHandler(h).
		WithGraphQLPlayground().
		WithMetrics(promhttp.HandlerFor(gath, promhttp.HandlerOpts{})).
		Run(context.Background(), addr)
}

func params() crmparams.Params {
	return crmparams.Params{
		SendPinByEmail:     GetEnvBool(EnvKeySendPinByEmail),
		ProjectPageBaseURL: GetEnvProjectPageBaseURL(),
	}
}
