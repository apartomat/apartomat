package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/apartomat/apartomat/api/crm/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/auth/paseto"
	"github.com/apartomat/apartomat/internal/dataloaders"
	"github.com/apartomat/apartomat/internal/image/minio"
	"github.com/apartomat/apartomat/internal/mail"
	"github.com/apartomat/apartomat/internal/mail/smtp"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	postgreshook "github.com/apartomat/apartomat/internal/postgres"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files/postgres"
	albums "github.com/apartomat/apartomat/internal/store/albums/postgres"
	contacts "github.com/apartomat/apartomat/internal/store/contacts/postgres"
	files "github.com/apartomat/apartomat/internal/store/files/postgres"
	houses "github.com/apartomat/apartomat/internal/store/houses/postgres"
	projects "github.com/apartomat/apartomat/internal/store/projects/postgres"
	sites "github.com/apartomat/apartomat/internal/store/public_sites/postgres"
	rooms "github.com/apartomat/apartomat/internal/store/rooms/postgres"
	users "github.com/apartomat/apartomat/internal/store/users/postgres"
	visualizations "github.com/apartomat/apartomat/internal/store/visualizations/postgres"
	workspaceUsers "github.com/apartomat/apartomat/internal/store/workspace_users/postgres"
	workspaces "github.com/apartomat/apartomat/internal/store/workspaces/postgres"
)

func main() {
	var (
		logLevel, _ = logLevel(os.Getenv("LOG_LEVEL"))
	)

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel, AddSource: false})))

	if len(os.Args) < 2 {
		slog.Error("expect command (run or gen-key-pair)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gen-key-pair":
		_, _, err := genPairToFile("apartomat.key")
		if err != nil {
			slog.Error("can't generate key pair", slog.Any("err", err))
			os.Exit(1)
		}

		slog.Info("done")

		os.Exit(0)

	case "run":
		privateKey, err := readPrivateKeyFromFile("apartomat.key")
		if err != nil {
			slog.Error("cant read private key from file", slog.Any("err", err))
			os.Exit(1)
		}

		confirmLoginIssuerVerifier := paseto.NewConfirmEmailTokenIssuerVerifier(privateKey)
		authIssuerVerifier := paseto.NewAuthTokenIssuerVerifier(privateKey)
		confirmEmailPin := paseto.NewConfirmEmailPINTokenIssuerVerifier(privateKey)
		invite := paseto.NewInviteTokenIssuerVerifier(privateKey)

		//
		reg := prometheus.NewRegistry()

		var (
			sqlHistrogramVec = prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name:    "sql_query_duration_seconds",
					Help:    "",
					Buckets: []float64{0.10, 0.2, 0.25, 0.3, 0.5, 1, 2, 2.5, 3, 5, 10},
				},
				[]string{"query"},
			)
		)

		reg.MustRegister(sqlHistrogramVec)

		//

		mailer := smtp.NewMailSender(smtp.Config{
			Addr:     os.Getenv("SMTP_ADDR"),
			User:     os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASSWORD"),
		})

		pgopts, err := pg.ParseURL(os.Getenv("POSTGRES_DSN"))
		if err != nil {
			slog.Error("can't parse POSTGRES_DSN", slog.Any("err", err))
			os.Exit(1)
		}

		pgdb := pg.Connect(pgopts)
		pgdb.AddQueryHook(postgreshook.NewLogQueryHook(slog.Default()))
		pgdb.AddQueryHook(postgreshook.NewQueryLatencyHook(func(dur time.Duration, query string) {
			sqlHistrogramVec.WithLabelValues(query).Observe(dur.Seconds())
		}))

		if err := pgdb.Ping(postgreshook.WithQueryContext(context.Background(), "ping")); err != nil {
			slog.Error("can't connect to database", slog.Any("err", err))
			os.Exit(1)
		}

		//

		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_DSN"))))

		bundb := bun.NewDB(sqldb, pgdialect.New())

		if err := bundb.Ping(); err != nil {
			slog.Error("can't connect to database", slog.Any("err", err))
			os.Exit(1)
		}

		bundb.AddQueryHook(bunhook.NewLogQueryHook(slog.Default()))
		bundb.AddQueryHook(bunhook.NewQueryLatencyHook(func(dur time.Duration, query string) {
			sqlHistrogramVec.WithLabelValues(query).Observe(dur.Seconds())
		}))

		//

		albumFilesStore := albumFiles.NewStore(bundb)
		albumsStore := albums.NewStore(bundb)
		contactsStore := contacts.NewStore(pgdb)
		filesStore := files.NewStore(bundb)
		housesStore := houses.NewStore(pgdb)
		projectsStore := projects.NewStore(bundb)
		publicSitesStore := sites.NewStore(bundb)
		roomsStore := rooms.NewStore(bundb)
		usersStore := users.NewStore(pgdb)
		visualizationsStore := visualizations.NewStore(bundb)
		workspaceUsersStore := workspaceUsers.NewStore(bundb)
		workspacesStore := workspaces.NewStore(pgdb)

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

		uploader := minio.NewUploader("apartomat")

		usecases := &apartomat.Apartomat{
			Params: apartomat.Params{
				SendPinByEmail: getBoolEnv("SEND_PIN_BY_EMAIL"),
			},
			AuthTokenIssuer:              authIssuerVerifier,
			AuthTokenVerifier:            authIssuerVerifier,
			ConfirmTokenByEmailIssuer:    confirmLoginIssuerVerifier,
			ConfirmTokenByEmailVerifier:  confirmLoginIssuerVerifier,
			ConfirmEmailPINTokenIssuer:   confirmEmailPin,
			ConfirmEmailPINTokenVerifier: confirmEmailPin,
			InviteTokenIssuer:            invite,
			InviteTokenVerifier:          invite,
			Mailer:                       mailer,
			MailFactory: mail.NewFactory(
				os.Getenv("BASE_URL"),
				os.Getenv("MAIL_FROM"),
			),
			Uploader:       uploader,
			Albums:         albumsStore,
			AlbumFiles:     albumFilesStore,
			Contacts:       contactsStore,
			Houses:         housesStore,
			Projects:       projectsStore,
			PublicSites:    publicSitesStore,
			Files:          filesStore,
			Rooms:          roomsStore,
			Users:          usersStore,
			Visualizations: visualizationsStore,
			Workspaces:     workspacesStore,
			WorkspaceUsers: workspaceUsersStore,
			Acl: apartomat.NewAcl(
				workspaceUsersStore,
				projectsStore,
				housesStore,
			),
		}

		var (
			addr = WithAddr(fmt.Sprintf(":%s", os.Getenv("PORT")))
		)

		if os.Getenv("SERVER_ADDR") != "" {
			addr = WithAddr(os.Getenv("SERVER_ADDR"))
		}

		h := graphql.Handler(
			usecases.CheckAuthToken,
			func() *dataloaders.DataLoaders {
				return dataloaders.NewLoaders(
					filesStore,
					roomsStore,
					usersStore,
					workspacesStore,
				)
			},
			graphql.NewRootResolver(bundb, usecases),
			10000,
		)

		NewServer(reg).WithGraphQLHandler(h).WithGraphQLPlayground().Run(context.Background(), addr)

	default:
		slog.Info("expect command (run or gen-key-pair)")
		os.Exit(1)
	}
}

func getBoolEnv(key string) bool {
	if val, err := strconv.ParseBool(os.Getenv(key)); err != nil {
		return val
	}

	return false
}
