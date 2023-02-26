package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) AddVisualizationsToAlbum(
	ctx context.Context,
	albumID string,
	visualizations []string,
) (AddVisualizationsToAlbumResult, error) {
	pages, err := r.useCases.AddVisualizationsToAlbum(ctx, albumID, visualizations)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't add visualization to album (id=%s): %s", albumID, err)

		return ServerError{Message: "can't add visualization to album"}, nil
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
