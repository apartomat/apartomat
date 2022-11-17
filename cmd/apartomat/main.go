package main

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/dataloader"
	"github.com/apartomat/apartomat/internal/image/minio"
	"github.com/apartomat/apartomat/internal/mail"
	"github.com/apartomat/apartomat/internal/mail/smtp"
	"github.com/apartomat/apartomat/internal/postgres/store"
	projects "github.com/apartomat/apartomat/internal/store/projects/postgres"
	visualizations "github.com/apartomat/apartomat/internal/store/visualizations/postgres"
	"github.com/apartomat/apartomat/internal/token"
	"github.com/go-pg/pg/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Print("expect command (run or gen-key-pair)")
		os.Exit(1)
	}

	ctx := context.Background()

	switch os.Args[1] {
	case "gen-key-pair":
		_, _, err := genPairToFile("shoppinglist.key")
		if err != nil {
			log.Fatalf("can't generate key pair: %s\n", err)
		}

		log.Print("done")
		os.Exit(0)

	case "run":
		privateKey, err := readPrivateKeyFromFile("shoppinglist.key")
		if err != nil {
			log.Fatalf("cant read private key from file: %s", err)
		}

		confirmLoginIssuerVerifier := token.NewPasetoMailConfirmTokenIssuerVerifier(privateKey)
		authIssuerVerifier := token.NewPasetoAuthTokenIssuerVerifier(privateKey)
		confirmEmailPin := token.NewPasetoConfirmEmailPINTokenIssuerVerifier(privateKey)

		mailer := smtp.NewMailSender(smtp.Config{
			Addr:     os.Getenv("SMTP_ADDR"),
			User:     os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASSWORD"),
		})

		pool, err := pgxpool.Connect(ctx, os.Getenv("POSTGRES_DSN"))
		if err != nil {
			log.Fatalf("can't connect to postgres: %s", err)
		}

		pgopts, err := pg.ParseURL(os.Getenv("POSTGRES_DSN"))
		if err != nil {
			log.Fatalf("can't parse POSTGRES_DSN %s", err)
		}

		pgdb := pg.Connect(pgopts)

		logger, err := NewLogger("debug")
		if err != nil {
			log.Fatalf("can't init zap: %s", err)
		}

		pgdb.AddQueryHook(loggerHook{"postgres", logger})

		users := store.NewUserStore(pool)
		workspaces := store.NewWorkspaceStore(pool)
		workspaceUsers := store.NewWorkspaceUserStore(pool)
		projectsStore := projects.NewStore(pgdb)
		projectFiles := store.NewProjectFileStore(pool)
		contactsStore := store.NewContactsStore(pgdb)
		housesStore := store.NewHousesStore(pgdb)
		roomsStore := store.NewRoomsStore(pgdb)
		visualizationsStore := visualizations.NewStore(pgdb)

		usersLoader := dataloader.NewUserLoader(dataloader.NewUserLoaderConfig(ctx, users))

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
			Mailer:                       mailer,
			MailFactory:                  mail.NewFactory("https://crm.apartomat.ru", "apartomat@yandex.ru"),
			Uploader:                     uploader,
			Contacts:                     contactsStore,
			Houses:                       housesStore,
			Projects:                     projectsStore,
			ProjectFiles:                 projectFiles,
			Rooms:                        roomsStore,
			Users:                        users,
			Visualizations:               visualizationsStore,
			Workspaces:                   workspaces,
			WorkspaceUsers:               workspaceUsers,
		}

		serverOpts := []Option{
			Addr(os.Getenv("SERVER_ADDR")),
		}

		NewServer(
			usecases,
			&dataloader.DataLoaders{
				Users: usersLoader,
			},
		).Run(serverOpts...)

	default:
		log.Print("expect command (run or gen-key-pair)")
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
