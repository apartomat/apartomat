package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/albums"
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
		slog.ErrorContext(ctx, "can't resolve project albums list", slog.String("err", "unknown project"))

		return serverError()
	} else {
		items, err := r.useCases.GetAlbums(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(
				ctx,
				"can't resolve project albums list",
				slog.String("project", project.ID),
				slog.Any("err", err),
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
		slog.ErrorContext(ctx, "can't resolve project albums total", slog.Any("err", errors.New("unknown project")))

		return serverError()
	} else {
		tot, err := r.useCases.CountAlbums(
			ctx,
			project.ID,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(
				ctx,
				"can't resolve project albums total",
				slog.String("project", project.ID),
				slog.Any("err", err),
			)

			return serverError()
		}

		return ProjectAlbumsTotal{Total: tot}, nil
	}
}
