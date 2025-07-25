package graphql

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/apartomat/apartomat/api/crm/graphql/dataloaders"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/albumfiles"
	"github.com/apartomat/apartomat/internal/store/albums"
)

func (r *rootResolver) Album() AlbumResolver { return &albumResolver{r} }

type albumResolver struct {
	*rootResolver
}

func (r *albumResolver) Pages(ctx context.Context, obj *Album) (AlbumPagesResult, error) {
	return obj.Pages, nil
}

func (r *albumResolver) Project(ctx context.Context, obj *Album) (AlbumProjectResult, error) {
	var (
		gp *Project
	)

	if pr, ok := obj.Project.(Project); ok {
		gp = &pr
	}

	if gp == nil {
		return serverError()
	}

	project, err := r.crm.GetProject(ctx, gp.ID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve project", slog.String("project", gp.ID), slog.Any("err", err))

		return nil, fmt.Errorf("can't resolve project (id=%s): %w", gp.ID, err)
	}

	return projectToGraphQL(project), nil
}

func (r *albumResolver) File(ctx context.Context, obj *Album) (AlbumRecentFileResult, error) {
	albumFile, file, err := r.crm.GetAlbumRecentFile(ctx, obj.ID)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, albums.ErrAlbumNotFound) {
			return notFound()
		}

		if errors.Is(err, albumfiles.ErrAlbumFileNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve recent album file", slog.String("project", obj.ID), slog.Any("err", err))

		return nil, fmt.Errorf("can't resolve album (id=%s) recent file: %w", obj.ID, err)
	}

	var (
		res = albumFileToGraphQL(albumFile)
	)

	if file != nil {
		res.File = fileToGraphQL(file)
	}

	return res, nil
}

func (r *albumResolver) Cover(ctx context.Context, obj *Album) (AlbumCoverResult, error) {
	if pages, ok := obj.Pages.(*AlbumPages); ok {
		if len(pages.Items) == 0 {
			return notFound()
		}

		if p, ok := pages.Items[0].(*AlbumPageCover); ok {
			if c, ok := p.Cover.(*CoverUploaded); ok {
				if f, ok := c.File.(File); ok {
					file, err := dataloaders.FromContext(ctx).Files.Load(ctx, f.ID)
					if err != nil {
						return nil, err
					}

					return fileToGraphQL(file), nil
				}
			}
		}

		return notFound()
	}

	slog.ErrorContext(ctx, "obj.Pages is not a *AlbumPages")

	return serverError()
}

func albumFileToGraphQL(file *albumfiles.AlbumFile) *AlbumFile {
	return &AlbumFile{
		ID:                  file.ID,
		Status:              albumFileStatusToGraphQL(file.Status),
		Version:             file.Version,
		GeneratingStartedAt: file.GeneratingStartedAt,
		GeneratingDoneAt:    file.GeneratingDoneAt,
	}
}

func albumFileStatusToGraphQL(status albumfiles.Status) AlbumFileStatus {
	switch status {
	case albumfiles.StatusNew:
		return AlbumFileStatusNew
	case albumfiles.StatusInProgress:
		return AlbumFileStatusGeneratingInProgress
	case albumfiles.StatusDone:
		return AlbumFileStatusGeneratingDone
	}

	return AlbumFileStatusNew
}
