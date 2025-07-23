package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) GenerateAlbumFile(ctx context.Context, albumID string) (GenerateAlbumFileResult, error) {
	albumFile, file, err := r.crm.StartGenerateAlbumFile(ctx, albumID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't start generate album file", slog.Any("err", err))

		return serverError()
	}

	var (
		res = albumFileToGraphQL(albumFile)
	)

	if file != nil {
		res.File = fileToGraphQL(file)
	}

	return AlbumFileGenerationStarted{File: res}, nil
}
