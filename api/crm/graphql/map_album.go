package graphql

import "github.com/apartomat/apartomat/internal/store/albums"

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
			Items: albumPagesToGraphQL(album.Pages, 0),
		},
	}
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
