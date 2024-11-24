package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/apartomat/apartomat/internal/project-page"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files/postgres"
	albums "github.com/apartomat/apartomat/internal/store/albums/postgres"
	files "github.com/apartomat/apartomat/internal/store/files/postgres"
	houses "github.com/apartomat/apartomat/internal/store/houses/postgres"
	sites "github.com/apartomat/apartomat/internal/store/public_sites/postgres"
	visualizations "github.com/apartomat/apartomat/internal/store/visualizations/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

func main() {
	var (
		addr = WithAddr(fmt.Sprintf(":%s", os.Getenv("PORT")))
	)

	if os.Getenv("SERVER_ADDR") != "" {
		addr = WithAddr(os.Getenv("SERVER_ADDR"))
	}

	ctx := context.Background()

	conf, err := pgx.ParseConfig(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(ctx, conf.ConnString())
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	var (
		bundb = bun.NewDB(
			sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_DSN")))),
			pgdialect.New(),
		)

		filesStore          = files.NewStore(bundb)
		publicSitesStore    = sites.NewStore(bundb)
		visualizationsStore = visualizations.NewStore(bundb)
		albumsStore         = albums.NewStore(bundb)
		albumsFilesStore    = albumFiles.NewStore(bundb)
		housesStore         = houses.NewStore(bundb)
	)

	publicSitesService := project_page.NewService(
		filesStore,
		publicSitesStore,
		visualizationsStore,
		albumsStore,
		albumsFilesStore,
		housesStore,
	)

	NewServer(publicSitesService).Run(addr)
}
