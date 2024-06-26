package main

import (
	"database/sql"
	"fmt"
	"github.com/apartomat/apartomat/internal/dataloaders"
	"log/slog"
	"os"
	"strconv"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/auth/paseto"
	"github.com/apartomat/apartomat/internal/image/minio"
	"github.com/apartomat/apartomat/internal/mail"
	"github.com/apartomat/apartomat/internal/mail/smtp"
	zapbun "github.com/apartomat/apartomat/internal/pkg/bun"
	"github.com/apartomat/apartomat/internal/postgres"
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
	workspace_users "github.com/apartomat/apartomat/internal/store/workspace_users/postgres"
	workspaces "github.com/apartomat/apartomat/internal/store/workspaces/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	var (
		logLevel, _ = logLevel(os.Getenv("LOG_LEVEL"))
	)

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})))

	// deprecated
	log, err := NewLogger(os.Getenv("LOG_LEVEL"))
	if err != nil {
		panic(err)
	}

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

		pgdb.AddQueryHook(postgres.NewZapLogQueryHook(log))

		reg := prometheus.NewRegistry()

		pgdb.AddQueryHook(postgres.NewQueryLatencyHook(reg))

		//

		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_DSN"))))

		bundb := bun.NewDB(sqldb, pgdialect.New())

		// todo: write to logger
		bundb.AddQueryHook(zapbun.NewZapLoggerQueryHook(log))

		bundb.AddQueryHook(zapbun.NewQueryLatencyHook(reg))

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
		workspaceUsersStore := workspace_users.NewStore(bundb)
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
			Logger:         log,
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

		NewServer(
			bundb,
			usecases,
			func() *dataloaders.DataLoaders {
				return dataloaders.NewLoaders(
					filesStore,
					roomsStore,
					usersStore,
					workspacesStore,
				)
			},
			reg,
			log,
		).Run(addr)

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
