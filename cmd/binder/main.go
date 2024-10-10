package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/crm/image/minio"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/apartomat/apartomat/internal/bookbinder"
	bunhook "github.com/apartomat/apartomat/internal/pkg/bun"
	. "github.com/apartomat/apartomat/internal/store/album_files"
	albumFilesPostgres "github.com/apartomat/apartomat/internal/store/album_files/postgres"
	"github.com/apartomat/apartomat/internal/store/albums"
	albumsPostgres "github.com/apartomat/apartomat/internal/store/albums/postgres"
	"github.com/apartomat/apartomat/internal/store/files"
	filesPostgres "github.com/apartomat/apartomat/internal/store/files/postgres"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	var (
		ctx = context.Background()

		db = bun.NewDB(
			sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_DSN")))),
			pgdialect.New(),
		)

		uploader = minio.NewUploader("apartomat")

		albumFilesStore = albumFilesPostgres.NewStore(db)
		albumsStore     = albumsPostgres.NewStore(db)
		filesStore      = filesPostgres.NewStore(db)
	)

	db.AddQueryHook(bunhook.NewLogQueryHook(slog.Default()))

	for {
		listFiles, err := albumFilesStore.List(ctx, And(StatusIn(StatusNew)), SortVersionAsc, 1, 0)
		if err != nil {
			slog.Error("can't list files", slog.Any("err", err))
			os.Exit(1)
		}

		slog.Info("files", slog.Int("len", len(listFiles)))

		for _, f := range listFiles {
			if err := f.StartNow(); err != nil {
				slog.Error("can't start generate file", slog.String("id", f.ID), slog.Any("err", err))
				os.Exit(1)
			}

			if err := albumFilesStore.Save(ctx, f); err != nil {
				slog.Error("can't save file", slog.String("id", f.ID), slog.Any("err", err))
				os.Exit(1)
			}

			slog.Info("start generate file", slog.String("id", f.ID))

			album, err := albumsStore.Get(ctx, albums.IDIn(f.AlbumID))
			if err != nil {
				slog.Error("can't find album", slog.String("id", f.AlbumID))
				os.Exit(1)
			}

			var (
				orient = bookbinder.Orientation(album.Settings.PageOrientation)
				format = bookbinder.Format(album.Settings.PageSize)
			)

			slog.Debug(
				"found album",
				slog.String("id", album.ID),
				slog.String("orientation", orient.String()),
				slog.String("format", format.String()),
				slog.Int("pages", len(album.Pages)),
			)

			var (
				pages = make([]string, len(album.Pages))
			)

			for i, p := range album.Pages {
				pages[i] = p.VisualizationID
			}

			images, err := download(filesStore)(ctx, album.Pages)
			if err != nil {
				slog.Error("can't download images", slog.Any("err", err))
				return
			}

			var (
				fileExt      = ".pdf"
				fileMimeType = mime.TypeByExtension(fileExt)
				fileName     = strings.Join([]string{f.ID, ".pdf"}, "")
				filePath     = fmt.Sprintf("pdf/%s/%s", album.ProjectID, fileName)
			)

			if r, err := bookbinder.Bind(orient, format, pages, images); err != nil {
				slog.Error("can't bind album", slog.String("id", album.ID), slog.Any("err", err), slog.String("mime", fileMimeType))
				os.Exit(1)
			} else {
				var (
					buf = &bytes.Buffer{}
					cp  = io.TeeReader(r, buf)
				)

				if b, err := io.ReadAll(cp); err != nil {
					slog.Error("can't read", slog.Any("err", err))
					os.Exit(1)
				} else {
					var (
						fileUrl = ""
					)

					if u, err := uploader.Upload(ctx, buf, int64(len(b)), filePath, mime.TypeByExtension(fileExt)); err != nil {
						slog.Error("can't upload file", slog.Any("err", err), slog.String("filePath", filePath))
						os.Exit(1)
					} else {
						slog.Info("uploaded", slog.String("filePath", u))
						fileUrl = u
					}

					if err := f.DoneNow(); err != nil {
						slog.Error("can't done file")
						os.Exit(1)
					}

					id, err := apartomat.GenerateNanoID()
					if err != nil {
						slog.Error("can't generate nano id", slog.Any("err", err))
						os.Exit(1)
					}

					var (
						file = files.NewFile(id, fileName, fileUrl, files.FileTypeAlbum, fileMimeType, album.ProjectID)
					)

					f.FileID = &file.ID

					if err := filesStore.Save(ctx, file); err != nil {
						slog.Error("can't save file", slog.Any("err", err))
						os.Exit(1)
					}

					if err := albumFilesStore.Save(ctx, f); err != nil {
						slog.Error("can't save album file", slog.String("id", f.ID), slog.Any("err", err))
						os.Exit(1)
					}
				}
			}
		}

		time.Sleep(3 * time.Second)
	}
}

func download(filesStore files.Store) func(ctx context.Context, pages []albums.AlbumPageVisualization) (map[string]io.Reader, error) {
	return func(ctx context.Context, pages []albums.AlbumPageVisualization) (map[string]io.Reader, error) {
		var (
			images = make(map[string]io.Reader, len(pages))

			ids = make([]string, len(pages))
		)

		for i, p := range pages {
			ids[i] = p.FileID
		}

		res, err := filesStore.List(ctx, files.IDIn(ids...), files.SortDefault, len(ids), 0)
		if err != nil {
			return nil, err
		}

		for i, f := range res {
			resp, err := http.Get(f.URL)
			if err != nil {
				return nil, err
			}

			if b, err := io.ReadAll(resp.Body); err != nil {
				return nil, err
			} else {
				images[pages[i].VisualizationID] = bytes.NewBuffer(b)
			}

			resp.Body.Close()
		}

		return images, nil
	}
}
