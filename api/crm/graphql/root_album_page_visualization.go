package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	svg "github.com/apartomat/apartomat/internal/crm/album"
)

func (r *rootResolver) AlbumPageVisualization() AlbumPageVisualizationResolver {
	return &albumPageVisualizationResolver{r}
}

type albumPageVisualizationResolver struct {
	*rootResolver
}

func (r *albumPageVisualizationResolver) SVG(ctx context.Context, obj *AlbumPageVisualization) (AlbumPageSVGResult, error) {
	var (
		album *Album
	)

	if a, ok := graphql.GetFieldContext(ctx).Parent.Parent.Parent.Parent.Result.(*Album); ok {
		album = a
	} else {
		slog.ErrorContext(ctx, "can't get album for visualization page")

		return serverError()
	}

	if v, ok := obj.Visualization.(*Visualization); ok && v != nil {
		vis, err := r.crm.GetVisualization(ctx, v.ID)
		if err != nil {
			if errors.Is(err, crm.ErrNotFound) {
				return notFound()
			}

			slog.ErrorContext(ctx, "can't get album page visualization", slog.Any("err", err))

			return serverError()
		}

		f, err := r.crm.GetFile(ctx, vis.FileID)
		if err != nil {
			if errors.Is(err, crm.ErrNotFound) {
				return notFound()
			}

			slog.ErrorContext(ctx, "can't get album page visualization file", slog.Any("err", err))

			return serverError()
		}

		res, err := svg.Visualization(
			graphQLToPageSize(album.Settings.PageSize),
			graphQLToPageOrientation(album.Settings.PageOrientation),
		)(obj.Number, f.URL)
		if err != nil {
			slog.ErrorContext(ctx, "can't get album page visualization svg", slog.Any("err", err))

			return serverError()
		}

		return SVG{SVG: res}, nil
	}

	slog.ErrorContext(ctx, "can't convert AlbumPageVisualization to Visualization")

	return serverError()
}

func (r *albumPageVisualizationResolver) Visualization(
	ctx context.Context,
	obj *AlbumPageVisualization,
) (AlbumPageVisualizationResult, error) {
	if v, ok := obj.Visualization.(*Visualization); ok && v != nil {
		vis, err := r.crm.GetVisualization(ctx, v.ID)
		if err != nil {
			if errors.Is(err, crm.ErrNotFound) {
				return notFound()
			}

			slog.ErrorContext(ctx, "can't get album page visualization", slog.Any("err", err))

			return serverError()
		}

		return visualizationToGraphQL(vis, nil), nil
	}

	slog.ErrorContext(ctx, "can't convert AlbumPageVisualization to Visualization")

	return serverError()
}
