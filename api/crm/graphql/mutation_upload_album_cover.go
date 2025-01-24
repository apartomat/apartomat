package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"log/slog"
)

func (r *mutationResolver) UploadAlbumCover(
	ctx context.Context,
	albumID string,
	file graphql.Upload,
) (UploadAlbumCoverResult, error) {
	uploaded, err := r.useCases.UploadAlbumCover(
		ctx,
		albumID,
		crm.Upload{
			Name:     file.Filename,
			MimeType: file.ContentType,
			Data:     file.File,
			Size:     file.Size,
		},
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(
			ctx,
			"can't upload file to album",
			slog.String("album", albumID),
			slog.Any("err", err),
		)

		return serverError()
	}

	return AlbumCoverUploaded{Cover: &CoverUploaded{File: fileToGraphQL(uploaded)}}, nil
}
