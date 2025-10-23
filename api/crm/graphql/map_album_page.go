package graphql

import "github.com/apartomat/apartomat/internal/store/albums"

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

func albumPageToGraphQL(p albums.AlbumPage, pageNumber int) AlbumPage {
	switch p.(type) {
	case albums.AlbumPageSplitCover:
		var (
			page = p.(albums.AlbumPageSplitCover)
		)

		return albumPageSplitCoverToGraphQL(page, pageNumber)
	case albums.AlbumPageCoverUploaded:
		var (
			page = p.(albums.AlbumPageCoverUploaded)
		)

		return albumPageCoverUploadedToGraphQL(page, pageNumber)
	case albums.AlbumPageVisualization:
		var (
			page = p.(albums.AlbumPageVisualization)
		)

		return albumPageVisualizationToGraphQL(page, pageNumber)
	}

	return nil
}

func albumPageSplitCoverToGraphQL(page albums.AlbumPageSplitCover, pageNumber int) *AlbumPageCover {
	imgSrc := ""
	qrCodeSrc := ""

	return &AlbumPageCover{
		ID:     page.ID,
		Number: pageNumber,
		Rotate: page.Rotate,
		Cover: &SplitCover{
			Title:     page.Title,
			Subtitle:  page.Subtitle,
			ImgSrc:    imgSrc,
			QRCodeSrc: &qrCodeSrc,
			City:      page.City,
			Year:      page.Year,
			Variant:   SplitCoverVariantImageOnTheRight,
		},
	}
}

func albumPageCoverUploadedToGraphQL(page albums.AlbumPageCoverUploaded, pageNumber int) *AlbumPageCover {
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
