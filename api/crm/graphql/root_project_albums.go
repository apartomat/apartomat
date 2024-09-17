package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"go.uber.org/zap"
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
		r.logger.Error("can't resolve project albums list", zap.Error(errors.New("unknown project")))

		return serverError()
	} else {
		items, err := r.useCases.GetAlbums(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			r.logger.Error(
				"can't resolve project albums list",
				zap.String("project", project.ID),
				zap.Error(err),
			)

			return serverError()
		}

		return ProjectAlbumsList{Items: albumsToGraphQL(items)}, nil
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
		r.logger.Error("can't resolve project albums total", zap.Error(errors.New("unknown project")))

		return serverError()
	} else {
		tot, err := r.useCases.CountAlbums(
			ctx,
			project.ID,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			r.logger.Error(
				"can't resolve project albums total",
				zap.String("project", project.ID),
				zap.Error(err),
			)

			return serverError()
		}

		return ProjectAlbumsTotal{Total: tot}, nil
	}
}
