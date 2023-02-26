package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"log"
)

func (r *mutationResolver) CreateAlbum(ctx context.Context, projectID, name string) (CreateAlbumResult, error) {
	album, err := r.useCases.CreateAlbum(
		ctx,
		projectID,
		name,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't create album: %s", err)

		return ServerError{Message: "can't create album"}, nil
	}

	return AlbumCreated{Album: albumToGraphQL(album)}, nil
}

func albumToGraphQL(album *albums.Album) *Album {
	return &Album{
		ID:   album.ID,
		Name: album.Name,
		Project: Project{
			ID: album.ProjectID,
		},
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
