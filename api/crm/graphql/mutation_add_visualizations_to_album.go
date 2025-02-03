package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) AddVisualizationsToAlbum(
	ctx context.Context,
	albumID string,
	visualizations []string,
) (AddVisualizationsToAlbumResult, error) {
	pages, n, err := r.useCases.AddVisualizationsToAlbum(ctx, albumID, visualizations)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(
			ctx,
			"can't add visualization to album",
			slog.String("album", albumID),
			slog.Any("err", err),
		)

		return serverError()
	}

	var (
		res = VisualizationsAddedToAlbum{
			Pages: make([]*AlbumPageVisualization, len(pages)),
		}

		num = n
	)

	for i, p := range pages {
		res.Pages[i] = albumPageVisualizationToGraphQL(p, num)
		num++
	}

	return res, nil
}
