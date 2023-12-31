package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"go.uber.org/zap"
)

func (r *mutationResolver) CreateAlbum(
	ctx context.Context,
	projectID, name string,
	settings CreateAlbumSettingsInput,
) (CreateAlbumResult, error) {
	album, err := r.useCases.CreateAlbum(
		ctx,
		projectID,
		name,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		r.logger.Error("can't create album", zap.Error(err))

		return serverError()
	}

	return AlbumCreated{Album: albumToGraphQL(album)}, nil
}

func albumToGraphQL(album *albums.Album) *Album {
	return &Album{
		ID:      album.ID,
		Name:    album.Name,
		Version: album.Version,
		Project: Project{
			ID: album.ProjectID,
		},
		Settings: albumSettingsToGraphQL(album.Settings),
		Pages: &AlbumPages{
			Items: albumPagesToGraphQL(album.Pages),
		},
	}
}

func albumPagesToGraphQL(pages []albums.AlbumPageVisualization) []AlbumPage {
	var (
		res = make([]AlbumPage, len(pages))
	)

	for i, p := range pages {
		res[i] = &AlbumPageVisualization{
			Position: 0,
			Visualization: &Visualization{
				ID: p.VisualizationID,
			},
		}
	}

	return res
}

func albumSettingsToGraphQL(settings albums.Settings) *AlbumSettings {
	var (
		res = &AlbumSettings{}
	)

	switch settings.PageSize {
	case albums.A4:
		res.PageSize = PageSizeA4
	case albums.A3:
		res.PageSize = PageSizeA3
	}

	switch settings.PageOrientation {
	case albums.Portrait:
		res.PageOrientation = PageOrientationPortrait
	case albums.Landscape:
		res.PageOrientation = PageOrientationLandscape
	}

	return res
}
