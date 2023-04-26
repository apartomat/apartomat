package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) AddVisualizationsToAlbum(
	ctx context.Context,
	albumID string,
	visualizations []string,
) (AddVisualizationsToAlbumResult, error) {
	pages, err := r.useCases.AddVisualizationsToAlbum(ctx, albumID, visualizations)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		r.logger.Error("can't add visualization to album", zap.String("album", albumID), zap.Error(err))

		return serverError()
	}

	var (
		res = VisualizationsAddedToAlbum{
			Pages: make([]*AlbumPageVisualization, len(pages)),
		}
	)

	for i, p := range pages {
		res.Pages[i] = &AlbumPageVisualization{
			Position:      p.Position,
			Visualization: visualizationToGraphQL(p.Visualization, nil),
		}
	}

	return res, nil
}
