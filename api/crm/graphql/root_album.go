package graphql

import (
	"context"
	"errors"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/dataloaders"
	albumFiles "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/apartomat/apartomat/internal/store/albums"
	"log/slog"
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

	project, err := r.useCases.GetProject(ctx, gp.ID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve project", slog.String("project", gp.ID), slog.Any("err", err))

		return nil, fmt.Errorf("can't resolve project (id=%s): %w", gp.ID, err)
	}

	return projectToGraphQL(project), nil
}

func (r *albumResolver) File(ctx context.Context, obj *Album) (AlbumRecentFileResult, error) {
	albumFile, file, err := r.useCases.GetAlbumRecentFile(ctx, obj.ID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, albums.ErrAlbumNotFound) {
			return notFound()
		}

		if errors.Is(err, albumFiles.ErrAlbumFileNotFound) {
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

func albumFileToGraphQL(file *albumFiles.AlbumFile) *AlbumFile {
	return &AlbumFile{
		ID:                  file.ID,
		Status:              albumFileStatusToGraphQL(file.Status),
		Version:             file.Version,
		GeneratingStartedAt: file.GeneratingStartedAt,
		GeneratingDoneAt:    file.GeneratingDoneAt,
	}
}

func albumFileStatusToGraphQL(status albumFiles.Status) AlbumFileStatus {
	switch status {
	case albumFiles.StatusNew:
		return AlbumFileStatusNew
	case albumFiles.StatusInProgress:
		return AlbumFileStatusGeneratingInProgress
	case albumFiles.StatusDone:
		return AlbumFileStatusGeneratingDone
	}

	return AlbumFileStatusNew
}
