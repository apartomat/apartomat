package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"log"
)

func (r *rootResolver) ProjectAlbums() ProjectAlbumsResolver {
	return &projectAlbumsResolver{r}
}

type projectAlbumsResolver struct {
	*rootResolver
}

func (r *projectAlbumsResolver) List(
	ctx context.Context,
	obj *ProjectAlbums,
	limit int,
	offset int,
) (ProjectAlbumsListResult, error) {
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		log.Printf("can't resolve project albums: %s", errors.New("unknown project"))

		return serverError()
	} else {
		albums, err := r.useCases.GetAlbums(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%s) albums: %s", project.ID, err)

			return ServerError{Message: "internal server error"}, nil
		}

		return ProjectAlbumsList{Items: albumsToGraphQL(albums)}, nil
	}
}

func albumsToGraphQL(albums []*albums.Album) []*Album {
	var (
		result = make([]*Album, 0, len(albums))
	)
	for _, u := range albums {
		result = append(result, albumToGraphQL(u))
	}

	return result
}

func (r *projectAlbumsResolver) Total(
	ctx context.Context,
	obj *ProjectAlbums,
) (ProjectAlbumsTotalResult, error) {
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		log.Printf("can't resolve project files: %s", errors.New("unknown project"))

		return nil, errors.New("server error: can't resolver project albums")
	} else {
		tot, err := r.useCases.CountAlbums(
			ctx,
			project.ID,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%s) files: %s", project.ID, err)

			return nil, errors.New("server error: can't resolver project albums")
		}

		return ProjectAlbumsTotal{Total: tot}, nil
	}
}
