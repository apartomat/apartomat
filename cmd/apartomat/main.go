package main

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/apartomat/apartomat/api/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/dataloader"
	"github.com/apartomat/apartomat/internal/image/s3"
	"github.com/apartomat/apartomat/internal/mail"
	"github.com/apartomat/apartomat/internal/postgres/store"
	"github.com/apartomat/apartomat/internal/token"
	"github.com/go-pg/pg/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"os"
)

var (
	serverOpts = []Option{
		Addr("localhost:8010"),
	}
)

func main() {
	if len(os.Args) == 0 {
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

		mailConfig := mail.SmtpServerConfig{
			Addr:     os.Getenv("SMTP_ADDR"),
			User:     os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASSWORD"),
		}

		confirmLoginIssuerVerifier := token.NewPasetoMailConfirmTokenIssuerVerifier(privateKey)
		authIssuerVerifier := token.NewPasetoAuthTokenIssuerVerifier(privateKey)
		mailer := mail.NewMailSender(mailConfig)

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
			log.Fatal("can't init zap: %s", err)
		}

		pgdb.AddQueryHook(loggerHook{"postgres", logger})

		users := store.NewUserStore(pool)
		workspaces := store.NewWorkspaceStore(pool)
		workspaceUsers := store.NewWorkspaceUserStore(pool)
		projects := store.NewProjectStore(pool)
		projectFiles := store.NewProjectFileStore(pool)
		contactsStore := store.NewContactsStore(pgdb)

		acl := apartomat.NewAcl(workspaceUsers)

		usersLoader := dataloader.NewUserLoader(dataloader.NewUserLoaderConfig(ctx, users))

		uploader, err := s3.NewS3ImageUploaderWithCred(
			ctx,
			os.Getenv("S3_ACCESS_KEY_ID"),
			os.Getenv("S3_SECRET_ACCESS_KEY"),
			os.Getenv("S3_REGION"),
			os.Getenv("S3_BUCKET_NAME"),
		)
		if err != nil {
			log.Fatalf("can't init s3: %s", err)
		}

		NewServer(
			&graphql.UseCases{
				CheckAuthToken:          apartomat.NewCheckAuthToken(authIssuerVerifier),
				LoginByEmail:            apartomat.NewLoginByEmail(users, workspaces, workspaceUsers, confirmLoginIssuerVerifier, mailer),
				ConfirmLogin:            apartomat.NewConfirmLogin(confirmLoginIssuerVerifier, authIssuerVerifier, users, acl),
				GetUserProfile:          apartomat.NewGetUserProfile(users),
				GetDefaultWorkspace:     apartomat.NewGetDefaultWorkspace(workspaces),
				GetWorkspace:            apartomat.NewGetWorkspace(workspaces, acl),
				GetWorkspaceUsers:       apartomat.NewGetWorkspaceUsers(workspaces, workspaceUsers, acl),
				GetWorkspaceUserProfile: apartomat.NewGetWorkspaceUserProfile(acl),
				GetWorkspaceProjects:    apartomat.NewGetWorkspaceProjects(workspaces, projects, acl),
				GetProject:              apartomat.NewGetProject(projects, acl),
				GetProjectFiles:         apartomat.NewGetProjectFiles(projects, projectFiles, acl),
				UploadProjectFile:       apartomat.NewUploadProjectFile(projects, projectFiles, acl, uploader),
				CreateProject:           apartomat.NewCreateProject(workspaces, projects, acl),
				GetContacts:             apartomat.NewGetContacts(projects, contactsStore, acl),
			},
			&apartomat.DataLoaders{
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
