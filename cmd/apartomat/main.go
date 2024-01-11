package main

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"go.uber.org/zap"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/auth/paseto"
	"github.com/apartomat/apartomat/internal/dataloader"
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
	log, err := NewLogger(os.Getenv("LOG_LEVEL"))
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("expect command (run or gen-key-pair)")
	}

	ctx := context.Background()

	switch os.Args[1] {
	case "gen-key-pair":
		_, _, err := genPairToFile("apartomat.key")
		if err != nil {
			log.Fatal("can't generate key pair", zap.Error(err))
		}

		log.Info("done")

		os.Exit(0)

	case "run":
		privateKey, err := readPrivateKeyFromFile("apartomat.key")
		if err != nil {
			log.Fatal("cant read private key from file", zap.Error(err))
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
			log.Fatal("can't parse POSTGRES_DSN %s", zap.Error(err))
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

		usersStore := users.NewStore(pgdb)
		workspacesStore := workspaces.NewStore(pgdb)
		workspaceUsersStore := workspace_users.NewStore(pgdb)
		projectsStore := projects.NewStore(pgdb)
		filesStore := files.NewStore(pgdb)
		albumsStore := albums.NewStore(pgdb)
		albumFilesStore := albumFiles.NewStore(bundb)
		contactsStore := contacts.NewStore(pgdb)
		housesStore := houses.NewStore(pgdb)
		roomsStore := rooms.NewStore(bundb)
		visualizationsStore := visualizations.NewStore(bundb)
		publicSitesStore := sites.NewStore(bundb)

		usersLoader := dataloader.NewUserLoader(dataloader.NewUserLoaderConfig(ctx, usersStore))

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
			&dataloader.DataLoaders{
				Users: usersLoader,
			},
			reg,
			log,
		).Run(addr)

	default:
		log.Info("expect command (run or gen-key-pair)")
		os.Exit(1)
	}
}

func readPrivateKeyFromFile(fileName string) (ed25519.PrivateKey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("can't read private key from file: %s", err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("can't read private key from file: %s", err)
	}

	block, _ := pem.Decode(b)

	bb, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %s", err)
	}

	return bb.(ed25519.PrivateKey), nil
}
