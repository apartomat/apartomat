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

func graphQLToPageSize(size PageSize) albums.PageSize {
	switch size {
	case PageSizeA3:
		return albums.A3
	case PageSizeA4:
		return albums.A4
	}

	return albums.A4
}

func graphQLToPageOrientation(orientation PageOrientation) albums.PageOrientation {
	switch orientation {
	case PageOrientationPortrait:
		return albums.Portrait
	case PageOrientationLandscape:
		return albums.Landscape
	}

	return albums.Portrait
}
