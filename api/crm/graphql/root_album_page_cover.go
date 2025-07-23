package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	svg "github.com/apartomat/apartomat/internal/crm/album"
)

func (r *rootResolver) AlbumPageCover() AlbumPageCoverResolver {
	return &albumPageCoverResolver{r}
}

type albumPageCoverResolver struct {
	*rootResolver
}

func (r *albumPageCoverResolver) SVG(ctx context.Context, obj *AlbumPageCover) (AlbumPageSVGResult, error) {
	var (
		album *Album
	)

	if a, ok := graphql.GetFieldContext(ctx).Parent.Parent.Parent.Parent.Result.(*Album); ok {
		album = a
	} else {
		slog.ErrorContext(ctx, "can't get album for cover")

		return serverError()
	}

	switch c := obj.Cover.(type) {
	case *CoverUploaded:
		if f, ok := c.File.(File); ok {
			file, err := r.crm.GetFile(ctx, f.ID)
			if err != nil {
				if errors.Is(err, crm.ErrNotFound) {
					return notFound()
				}

				slog.ErrorContext(ctx, "can't get album cover file", slog.Any("err", err))

				return serverError()
			}

			res, err := svg.UploadedCover(
				graphQLToPageSize(album.Settings.PageSize),
				graphQLToPageOrientation(album.Settings.PageOrientation),
			)(obj.Number, file.URL)
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
