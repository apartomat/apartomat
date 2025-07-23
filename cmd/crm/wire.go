//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"crypto/ed25519"
	"database/sql"
	"fmt"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/apartomat/apartomat/internal/crm/auth/paseto"
	"github.com/apartomat/apartomat/internal/crm/image"
	"github.com/apartomat/apartomat/internal/crm/image/minio"
	"github.com/apartomat/apartomat/internal/crm/mail"
	"github.com/apartomat/apartomat/internal/crm/mail/smtp"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	gopghook "github.com/apartomat/apartomat/internal/pkg/go-pg"
	albumfiles "github.com/apartomat/apartomat/internal/store/albumfiles/postgres"
	albums "github.com/apartomat/apartomat/internal/store/albums/postgres"
	contacts "github.com/apartomat/apartomat/internal/store/contacts/postgres"
	files "github.com/apartomat/apartomat/internal/store/files/postgres"
	houses "github.com/apartomat/apartomat/internal/store/houses/postgres"
	projectpage "github.com/apartomat/apartomat/internal/store/projectpages/postgres"
	projects "github.com/apartomat/apartomat/internal/store/projects/postgres"
	rooms "github.com/apartomat/apartomat/internal/store/rooms/postgres"
	users "github.com/apartomat/apartomat/internal/store/users/postgres"
	visualizations "github.com/apartomat/apartomat/internal/store/visualizations/postgres"
	workspaces "github.com/apartomat/apartomat/internal/store/workspaces/postgres"
	workspaceusers "github.com/apartomat/apartomat/internal/store/workspaceusers/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/google/wire"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log/slog"
	"os"
)

func InitializeCRM(ctx context.Context) (*crm.CRM, error) {
	wire.Build(
		crm.NewAcl,
		crm.NewCRM,
		params,

		ProvidePrivateKey,
		ProvideBun,
		ProvidePg,
		ProvideUploader,
		ProvideMailFactory,
		ProvideMailSender,

		StoreSet,

		paseto.NewAuthTokenIssuerVerifier,
		paseto.NewConfirmEmailTokenIssuerVerifier,
		paseto.NewConfirmEmailPINTokenIssuerVerifier,
		paseto.NewInviteTokenIssuerVerifier,

		wire.Bind(new(auth.AuthTokenIssuer), new(*paseto.AuthTokenIssuerVerifier)),
		wire.Bind(new(auth.AuthTokenVerifier), new(*paseto.AuthTokenIssuerVerifier)),

		wire.Bind(new(auth.EmailConfirmTokenIssuer), new(*paseto.ConfirmEmailTokenIssuerVerifier)),
		wire.Bind(new(auth.EmailConfirmTokenVerifier), new(*paseto.ConfirmEmailTokenIssuerVerifier)),

		wire.Bind(new(auth.ConfirmEmailPINTokenIssuer), new(*paseto.ConfirmEmailPINTokenIssuerVerifier)),
		wire.Bind(new(auth.ConfirmEmailPINTokenVerifier), new(*paseto.ConfirmEmailPINTokenIssuerVerifier)),

		wire.Bind(new(auth.InviteTokenIssuer), new(*paseto.InviteTokenIssuerVerifier)),
		wire.Bind(new(auth.InviteTokenVerifier), new(*paseto.InviteTokenIssuerVerifier)),
	)

	return &crm.CRM{}, nil
}

func ProvidePrivateKey() (ed25519.PrivateKey, error) {
	privateKey, err := readPrivateKeyFromFile("apartomat.key")
	if err != nil {
		return nil, fmt.Errorf("cant read private key from file: %w", err)
	}

	return privateKey, nil
}

func ProvideBun() (*bun.DB, error) {
	var (
		sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_DSN"))))
		bundb = bun.NewDB(sqldb, pgdialect.New())
	)

	if err := bundb.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	bundb.AddQueryHook(bunhook.NewLogQueryHook(slog.Default()))
	bundb.AddQueryHook(bunhook.NewQueryLatencyHook(observeSql))

	return bundb, nil
}

func ProvidePg(ctx context.Context) (*pg.DB, error) {
	pgopts, err := pg.ParseURL(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, fmt.Errorf("can't parse POSTGRES_DSN: %w", err)
	}

	db := pg.Connect(pgopts)

	db.AddQueryHook(gopghook.NewLogQueryHook(slog.Default()))
	db.AddQueryHook(gopghook.NewQueryLatencyHook(observeSql))

	if err := db.Ping(gopghook.WithQueryContext(ctx, "ping")); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return db, nil
}

func ProvideUploader() (image.Uploader, error) {
	return minio.NewUploader("apartomat"), nil
}

func ProvideMailFactory() (*mail.Factory, error) {
	return mail.NewFactory(
		os.Getenv("BASE_URL"),
		os.Getenv("MAIL_FROM"),
	), nil
}

func ProvideMailSender() (mail.Sender, error) {
	return smtp.NewMailSender(smtp.Config{
		Addr:     os.Getenv("SMTP_ADDR"),
		User:     os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}), nil
}

var StoreSet = wire.NewSet(
	albumfiles.Set,
	albums.Set,
	contacts.Set,
	files.Set,
	houses.Set,
	projectpage.Set,
	projects.Set,
	rooms.Set,
	users.Set,
	visualizations.Set,
	workspaces.Set,
	workspaceusers.Set,
)
