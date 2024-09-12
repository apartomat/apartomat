package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/crm/svg"
	"go.uber.org/zap"
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

			r.logger.Error("can't get album page visualization", zap.Error(err))

			return serverError()
		}

		f, err := r.useCases.GetFile(ctx, vis.FileID)
		if err != nil {
			if errors.Is(err, apartomat.ErrNotFound) {
				return notFound()
			}

			r.logger.Error("can't get album page visualization file", zap.Error(err))

			return serverError()
		}

		return SVG{
			SVG: svg.Visualization(obj.Number, f.URL),
		}, nil
	}

	r.logger.Error("can't convert AlbumPageVisualization to Visualization")

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

			r.logger.Error("can't get album page visualization", zap.Error(err))

			return serverError()
		}

		return visualizationToGraphQL(vis, nil), nil
	}

	r.logger.Error("can't convert AlbumPageVisualization to Visualization")

	return serverError()
}
