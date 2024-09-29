package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log/slog"
)

func (r *mutationResolver) GenerateAlbumFile(ctx context.Context, albumID string) (GenerateAlbumFileResult, error) {
	albumFile, file, err := r.useCases.StartGenerateAlbumFile(ctx, albumID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
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
