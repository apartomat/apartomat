package main

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ztsu/apartomat/api/graphql"
	"github.com/ztsu/apartomat/internal"
	"github.com/ztsu/apartomat/internal/store/postgres"
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

		mailConfig := apartomat.SmtpServerConfig{
			Addr:     os.Getenv("SMTP_ADDR"),
			User:     os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASSWORD"),
		}

		confirmLoginIssuerVerifier := apartomat.NewPasetoMailConfirmTokenIssuerVerifier(privateKey)
		authIssuerVerifier := apartomat.NewPasetoAuthTokenIssuerVerifier(privateKey)
		mailer := apartomat.NewMailSender(mailConfig)

		pool, err := pgxpool.Connect(ctx, os.Getenv("POSTGRES_DSN"))
		if err != nil {
			log.Fatalf("can't connect to postgres: %s", err)
		}

		users := postgres.NewUserStore(pool)
		workspaces := postgres.NewWorkspaceStore(pool)
		workspaceUsers := postgres.NewWorkspaceUserStore(pool)

		NewServer(
			&graphql.UseCases{
				CheckAuthToken:      apartomat.NewCheckAuthToken(authIssuerVerifier),
				LoginByEmail:        apartomat.NewLoginByEmail(users, workspaces, workspaceUsers, confirmLoginIssuerVerifier, mailer),
				ConfirmLogin:        apartomat.NewConfirmLogin(confirmLoginIssuerVerifier, authIssuerVerifier),
				GetUserProfile:      apartomat.NewGetUserProfile(users),
				GetDefaultWorkspace: apartomat.NewGetDefaultWorkspace(workspaces),
				GetWorkspace:        apartomat.NewGetWorkspace(workspaces),
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
