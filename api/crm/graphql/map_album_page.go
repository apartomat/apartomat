package graphql

import "github.com/apartomat/apartomat/internal/store/albums"

func albumPageToGraphQL(p albums.AlbumPage, pageNumber int) AlbumPage {
	switch p.(type) {
	case albums.AlbumPageCover:
		var (
			page = p.(albums.AlbumPageCover)
		)
		return &AlbumPageCover{
			ID:     page.ID,
			Number: pageNumber,
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

		return &AlbumPageCover{
			ID:     page.ID,
			Number: pageNumber,
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

		return albumPageVisualizationToGraphQL(page, pageNumber)
	}

	return nil
}

func albumPagesToGraphQL(pages []albums.AlbumPage, startPageNumber int) []AlbumPage {
	var (
		res = make([]AlbumPage, len(pages))
		num = startPageNumber
	)

	for i, p := range pages {
		res[i] = albumPageToGraphQL(p, num)
		num++
	}

	return res
}

func albumPageVisualizationToGraphQL(page albums.AlbumPageVisualization, pageNumber int) *AlbumPageVisualization {
	return &AlbumPageVisualization{
		ID:     page.ID,
		Number: pageNumber,
		Rotate: page.Rotate,
		Visualization: &Visualization{
			ID: page.VisualizationID,
		},
	}
}
