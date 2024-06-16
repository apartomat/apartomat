package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/bookbinder"
	"github.com/apartomat/apartomat/internal/image/minio"
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
	"go.uber.org/zap"
	"io"
	"mime"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	log, err := NewLogger("debug")
	if err != nil {
		panic(err)
	}

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

	db.AddQueryHook(bunhook.NewZapLoggerQueryHook(log))

	for {
		listFiles, err := albumFilesStore.List(ctx, And(StatusIn(StatusNew)), SortVersionAsc, 1, 0)
		if err != nil {
			log.Fatal("can't list files", zap.Error(err))
		}

		log.Info("files", zap.Int("len", len(listFiles)))

		for _, f := range listFiles {
			if err := f.StartNow(); err != nil {
				log.Fatal("can't start generate file", zap.String("id", f.ID), zap.Error(err))
			}

			if err := albumFilesStore.Save(ctx, f); err != nil {
				log.Fatal("can't save file", zap.String("id", f.ID), zap.Error(err))
			}

			log.Info("start generate file", zap.String("id", f.ID))

			album, err := albumsStore.Get(ctx, albums.IDIn(f.AlbumID))
			if err != nil {
				log.Fatal("can't find album", zap.String("id", f.AlbumID))
			}

			var (
				orient = bookbinder.Orientation(album.Settings.PageOrientation)
				format = bookbinder.Format(album.Settings.PageSize)
			)

			log.Debug(
				"found album",
				zap.String("id", album.ID),
				zap.String("orientation", orient.String()),
				zap.String("format", format.String()),
				zap.Int("pages", len(album.Pages)),
			)

			var (
				pages = make([]string, len(album.Pages))
			)

			for i, p := range album.Pages {
				pages[i] = p.VisualizationID
			}

			images, err := download(filesStore)(ctx, album.Pages)
			if err != nil {
				log.Error("can't download images", zap.Error(err))
				return
			}

			var (
				fileExt      = ".pdf"
				fileMimeType = mime.TypeByExtension(fileExt)
				fileName     = strings.Join([]string{f.ID, ".pdf"}, "")
				filePath     = fmt.Sprintf("pdf/%s/%s", album.ProjectID, fileName)
			)

			if r, err := bookbinder.Bind(orient, format, pages, images); err != nil {
				log.Fatal("can't bind album", zap.String("id", album.ID), zap.Error(err), zap.String("mime", fileMimeType))
			} else {
				var (
					buf = &bytes.Buffer{}
					cp  = io.TeeReader(r, buf)
				)

				if b, err := io.ReadAll(cp); err != nil {
					log.Fatal("can't read", zap.Error(err))
				} else {
					var (
						fileUrl = ""
					)

					if u, err := uploader.Upload(ctx, buf, int64(len(b)), filePath, mime.TypeByExtension(fileExt)); err != nil {
						log.Fatal("can't upload file", zap.Error(err), zap.String("filePath", filePath))
					} else {
						log.Info("uploaded", zap.String("filePath", u))
						fileUrl = u
					}

					if err := f.DoneNow(); err != nil {
						log.Fatal("can't done file")
					}

					id, err := apartomat.GenerateNanoID()
					if err != nil {
						log.Fatal("can't generate nano id", zap.Error(err))
					}

					var (
						file = files.NewFile(id, fileName, fileUrl, files.FileTypeAlbum, fileMimeType, album.ProjectID)
					)

					f.FileID = &file.ID

					if err := filesStore.Save(ctx, file); err != nil {
						log.Fatal("can't save file", zap.Error(err))
					}

					if err := albumFilesStore.Save(ctx, f); err != nil {
						log.Fatal("can't save album file", zap.String("id", f.ID), zap.Error(err))
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
