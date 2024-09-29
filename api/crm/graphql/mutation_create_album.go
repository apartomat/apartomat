package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"log/slog"
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

		slog.ErrorContext(ctx, "can't create album", slog.Any("err", err))

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

func albumPagesToGraphQL(pages []albums.AlbumPage) []AlbumPage {
	var (
		res = make([]AlbumPage, len(pages))
	)

	for i, p := range pages {
		switch p.(type) {
		case albums.AlbumPageCover:
			var (
				page = p.(albums.AlbumPageCover)
			)
			res[i] = &AlbumPageCover{
				Number: i,
				Rotate: page.Rotate,
				Cover: &CoverUploaded{
					File: File{
						ID: page.FileID,
					},
				},
			}
		case albums.AlbumPageCoverUploaded:
			var (
				page = p.(albums.AlbumPageCoverUploaded)
			)
			res[i] = &AlbumPageCover{
				Number: i,
				Rotate: page.Rotate,
				Cover: &CoverUploaded{
					File: File{
						ID: page.FileID,
					},
				},
			}
		case albums.AlbumPageVisualization:
			var (
				page = p.(albums.AlbumPageVisualization)
			)

			res[i] = &AlbumPageVisualization{
				Number: i,
				Rotate: page.Rotate,
				Visualization: &Visualization{
					ID: page.VisualizationID,
				},
			}
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
