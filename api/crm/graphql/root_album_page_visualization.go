package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/crm/svg"
)

func (r *rootResolver) AlbumPageVisualization() AlbumPageVisualizationResolver {
	return &albumPageVisualizationResolver{r}
}

type albumPageVisualizationResolver struct {
	*rootResolver
}

func (r *albumPageVisualizationResolver) SVG(ctx context.Context, obj *AlbumPageVisualization) (AlbumPageSVGResult, error) {
	if v, ok := obj.Visualization.(*Visualization); ok && v != nil {
		vis, err := r.useCases.GetVisualization(ctx, v.ID)
		if err != nil {
			if errors.Is(err, apartomat.ErrNotFound) {
				return notFound()
			}

			slog.ErrorContext(ctx, "can't get album page visualization", slog.Any("err", err))

			return serverError()
		}

		f, err := r.useCases.GetFile(ctx, vis.FileID)
		if err != nil {
			if errors.Is(err, apartomat.ErrNotFound) {
				return notFound()
			}

			slog.ErrorContext(ctx, "can't get album page visualization file", slog.Any("err", err))

			return serverError()
		}

		res, err := svg.Visualization(obj.Number, f.URL)
		if err != nil {
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
		vis, err := r.useCases.GetVisualization(ctx, v.ID)
		if err != nil {
			if errors.Is(err, apartomat.ErrNotFound) {
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
