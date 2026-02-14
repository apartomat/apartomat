package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"

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

		slog.ErrorContext(ctx, "can't resolve album project: CoverUploaded.File is not a File")

		return serverError()
	case *SplitCover:
		slog.ErrorContext(ctx, "can't get svg for split cover", slog.String("err", "not implemented yet"))
		return notImplementedYetError()
	default:
		slog.ErrorContext(ctx, "unknown obj type")
		return serverError()
	}
}
