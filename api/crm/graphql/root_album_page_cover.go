package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/crm/svg"
)

func (r *rootResolver) AlbumPageCover() AlbumPageCoverResolver {
	return &albumPageCoverResolver{r}
}

type albumPageCoverResolver struct {
	*rootResolver
}

func (r *albumPageCoverResolver) SVG(ctx context.Context, obj *AlbumPageCover) (AlbumPageSVGResult, error) {
	switch c := obj.Cover.(type) {
	case *CoverUploaded:
		if f, ok := c.File.(File); ok {
			file, err := r.useCases.GetFile(ctx, f.ID)
			if err != nil {
				if errors.Is(err, apartomat.ErrNotFound) {
					return notFound()
				}

				slog.ErrorContext(ctx, "can't get album cover file", slog.Any("err", err))

				return serverError()
			}

			res, err := svg.UploadedCover(obj.Number, file.URL)
			if err != nil {
				slog.ErrorContext(ctx, "can't get svg for uploaded cover", slog.Any("err", err))

				return serverError()
			}

			return SVG{res}, nil
		}

		slog.ErrorContext(ctx, "can't convert AlbumPageCover.File to CoverUploaded")

		return serverError()
	case *Cover:
		return notImplementedYetError()
	default:
		slog.ErrorContext(ctx, "unknown obj type")
		return serverError()
	}
}

func (r *albumPageCoverResolver) Cover(
	ctx context.Context,
	obj *AlbumPageCover,
) (AlbumPageCoverResult, error) {
	var (
		cov interface{} = obj.Cover
	)

	switch c := cov.(type) {
	case *CoverUploaded:
		return c, nil
	case *Cover:
		return notImplementedYetError()
	default:
		slog.ErrorContext(ctx, "unknown obj type")
		return serverError()
	}
}
