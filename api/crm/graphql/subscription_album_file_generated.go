package graphql

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/apartomat/apartomat/internal/crm"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/apartomat/apartomat/internal/store/albums"
)

func (r *subscriptionResolver) AlbumFileGenerated(ctx context.Context, id string) (<-chan AlbumFileGenerated, error) {
	var (
		ch = make(chan AlbumFileGenerated)
	)

	go func() {
		defer close(ch)

		var (
			tck = time.NewTicker(1 * time.Second)

			tr = time.NewTimer(60 * time.Second)
		)

		for {
			select {
			case <-ctx.Done():
				ch <- Unknown{"context closed"}
				return
			case <-tr.C:
				ch <- Unknown{"timeout"}
				return
			case <-tck.C:
				albumFile, _, err := r.useCases.GetAlbumRecentFile(ctx, id)
				if err != nil {
					switch {
					case errors.Is(err, crm.ErrForbidden):
						ch <- Forbidden{Message: "forbidden"}
					case errors.Is(err, albums.ErrAlbumNotFound):
						ch <- NotFound{Message: "not found"}
					case errors.Is(err, albumFiles.ErrAlbumFileNotFound):
						ch <- NotFound{Message: "not found"}
					default:
						ch <- ServerError{Message: "server error"}
						slog.ErrorContext(ctx, "can't resolve recent album file", slog.String("project", id), slog.Any("err", err))
					}

					return
				}

				if albumFile.Status == albumFiles.StatusDone {
					ch <- albumFileToGraphQL(albumFile)
				}
			}
		}
	}()

	return ch, nil
}
